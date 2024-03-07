package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Pallinder/sillyname-go"
	"go.uber.org/zap"

	"dimoklan/consts"
	"dimoklan/consts/hashtag"
	"dimoklan/consts/newuser"
	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/repo"
	"dimoklan/util"
)

type RegisterService struct {
	core        config.Core
	storage     repo.Storage
	cellService *CellService
}

func NewRegisterService(core config.Core, storage repo.Storage, cellService *CellService) *RegisterService {
	return &RegisterService{
		core:        core,
		storage:     storage,
		cellService: cellService,
	}
}

func (rs *RegisterService) Create(ctx context.Context, register model.Register) (model.Register, error) {
	if err := register.ValidateRegister(); err != nil {
		return model.Register{}, err
	}

	hashedEmail := consts.HashSalt + register.Email + rs.core.GetSalt()

	// Calculate the SHA-256 hash of the byte slice
	activationCode := sha256.Sum256([]byte(hashedEmail))
	activationCodeHashed := sha256.Sum256([]byte(hex.EncodeToString(activationCode[:])))

	register.ActivationCode = hex.EncodeToString(activationCodeHashed[:])
	// delete after 24 hours
	register.Language = consts.LanguageEn
	register.Password = util.HashPassword(register.Password, consts.HashSalt, rs.core.GetSalt())

	// check if user already registered with same email
	user, err := rs.storage.GetUserByEmail(ctx, register.Email)
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return register, err
	}

	if user.Email != "" {
		return register, errors.New("email is not avaialble")
	}

	if err := rs.storage.CreateRegister(ctx, register.ToRepo()); err != nil {
		rs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return register, err
	}

	// TODO: this should be sent by email
	fmt.Println(">>>> actiation code: ", hex.EncodeToString(activationCode[:]))
	register.Password = ""
	register.ActivationCode = ""
	return register, nil
}

func (rs *RegisterService) Confirm(ctx context.Context, activationCode string) error {
	activationCodeHashed := sha256.Sum256([]byte(activationCode))

	registerRepo, err := rs.storage.ConfirmRegister(ctx, hashtag.Register+hex.EncodeToString(activationCodeHashed[:]))
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("activation_failed"))
		return err
	}

	register := registerRepo.ToAPI()

	if register.Email == "" {
		return errors.New("activation is not valid")
	}

	// check if user already registered with same email
	tmpUser, err := rs.storage.GetUserByEmail(ctx, hashtag.User+register.Email)
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("activation_failed"))
		return err
	}

	if tmpUser.Email != "" {
		return errors.New("activation has already been completed")
	}

	id := rand.Intn(consts.MaxUserID)

	// create user
	user := model.User{
		ID:            strconv.Itoa(id),
		Color:         strconv.FormatInt(int64(id), 16),
		Farr:          newuser.Farr,
		Gold:          newuser.Gold,
		Email:         register.Email,
		Kingdom:       register.Kingdom,
		Password:      register.Password,
		Language:      register.Language,
		Suspend:       false,
		SuspendReason: "",
		Freeze:        false,
		FreezeReason:  "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := rs.storage.CreateUser(ctx, user.ToRepo()); err != nil {
		rs.core.Error(err.Error(), zap.Stack("user_creation_failed"))
		return err
	}

	// create auth for login
	auth := model.Auth{
		UserID:        user.ID,
		Email:         register.Email,
		Password:      register.Password,
		Suspend:       false,
		SuspendReason: "",
	}

	if err := rs.storage.CreateAuth(ctx, auth.ToRepo()); err != nil {
		rs.storage.DeleteUser(ctx, hashtag.User+user.ID)
		rs.core.Error(err.Error(), zap.Stack("auth_creation_failed"))
		return err
	}

	// create a marshal for user
	marshal := model.Marshal{
		UserID:    user.ID,
		ID:        user.ID + ":1",
		Name:      sillyname.GenerateStupidName(),
		Cell:      register.Cell,
		Army:      newuser.Army,
		Star:      newuser.Star,
		Speed:     newuser.Speed,
		Attack:    newuser.Attack,
		Face:      "todo_to_be_added",
		CreatedAt: time.Now(),
	}
	if err := rs.storage.CreateMarshal(ctx, marshal.ToRepo()); err != nil {
		rs.storage.DeleteUser(ctx, hashtag.User+user.ID)
		rs.storage.DeleteAuth(ctx, hashtag.Auth+auth.Email)
		rs.core.Error(err.Error(), zap.Stack("marshal_creation_failed"))
		return err
	}

	cell := model.Cell{
		Cell: register.Cell,
	}
	if err := rs.cellService.AssignCellToUser(ctx, cell, marshal.UserID); err != nil {
		rs.storage.DeleteUser(ctx, hashtag.User+user.ID)
		rs.storage.DeleteAuth(ctx, hashtag.Auth+auth.Email)
		rs.storage.DeleteMarshal(ctx, hashtag.User+user.ID, hashtag.Marshal+marshal.ID)
		rs.core.Error(err.Error(), zap.Stack("error_in_assigning_cell_to_user"))
		return err
	}

	return nil
}
