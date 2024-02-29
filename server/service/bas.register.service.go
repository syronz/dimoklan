package service

import (
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
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/types"
	"dimoklan/util"
)

type RegisterService struct {
	core        config.Core
	storage     basstorage.BasStorage
	cellService *CellService
}

func NewRegisterService(core config.Core, storage basstorage.BasStorage, cellService *CellService) *RegisterService {
	return &RegisterService{
		core:        core,
		storage:     storage,
		cellService: cellService,
	}
}

func (rs *RegisterService) Create(register types.Register) (types.Register, error) {
	if err := register.ValidateRegister(); err != nil {
		return types.Register{}, err
	}

	hashedEmail := consts.HashSalt + register.Email + rs.core.GetSalt()

	// Calculate the SHA-256 hash of the byte slice
	activationCode := sha256.Sum256([]byte(hashedEmail))
	activationCodeHashed := sha256.Sum256([]byte(hex.EncodeToString(activationCode[:])))

	register.ActivationCode = hex.EncodeToString(activationCodeHashed[:])
	// delete after 24 hours
	register.TTL = time.Now().Add(24 * time.Hour).Unix()
	register.Language = consts.LanguageEn
	register.Password = util.HashPassword(register.Password, consts.HashSalt, rs.core.GetSalt())

	// check if user already registered with same email
	user, err := rs.storage.GetUserByEmail(register.Email)
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return register, err
	}

	if user.Email != "" {
		return register, errors.New("email is not avaialble")
	}

	if err := rs.storage.CreateRegister(register); err != nil {
		rs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return register, err
	}

	// TODO: this should be sent by email
	fmt.Println(">>>> actiation code: ", hex.EncodeToString(activationCode[:]))
	register.Password = ""
	register.ActivationCode = ""
	return register, nil
}

func (rs *RegisterService) Confirm(activationCode string) error {
	activationCodeHashed := sha256.Sum256([]byte(activationCode))

	register, err := rs.storage.ConfirmRegister(hex.EncodeToString(activationCodeHashed[:]))
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("activation_failed"))
		return err
	}

	// check if user already registered with same email
	tmpUser, err := rs.storage.GetUserByEmail(register.Email)
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("activation_failed"))
		return err
	}

	if tmpUser.Email != "" {
		return errors.New("activation has already been completed")
	}

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(consts.MaxUserID)

	// create user
	user := types.User{
		ID:            strconv.Itoa(id),
		Color:         strconv.FormatInt(int64(id), 16),
		Farr:          consts.FarrForNewUser,
		Gold:          consts.GoldForNewUser,
		Email:         register.Email,
		Kingdom:       register.Kingdom,
		Password:      register.Password,
		Language:      register.Language,
		Suspend:       false,
		SuspendReason: "",
		Freeze:        false,
		FreezeReason:  "",
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
	}

	if err := rs.storage.CreateUser(user); err != nil {
		rs.core.Error(err.Error(), zap.Stack("user_creation_failed"))
		return err
	}

	// create auth for login
	auth := types.Auth{
		UserID:        user.ID,
		Email:         register.Email,
		Password:      register.Password,
		Suspend:       false,
		SuspendReason: "",
	}

	if err := rs.storage.CreateAuth(auth); err != nil {
		rs.storage.DeleteUser(consts.ParUser + user.ID)
		rs.core.Error(err.Error(), zap.Stack("auth_creation_failed"))
		return err
	}

	// create a marshal for user
	marshal := types.Marshal{
		UserID:     user.ID,
		ID:         user.ID + ":1",
		Name:       sillyname.GenerateStupidName(),
		Cell:       register.Cell,
		Army:       consts.ArmyForNewUser,
		Star:       consts.StarForNewUser,
		Speed:      consts.SpeedForNewUser,
		Attack:     consts.AttackForNewUser,
		Face:       "todo_to_be_added",
		CreatedAt:  time.Now().Unix(),
		EntityType: consts.MarshalEntity,
	}
	if err := rs.storage.CreateMarshal(marshal); err != nil {
		rs.storage.DeleteUser(consts.ParUser + user.ID)
		rs.storage.DeleteAuth(consts.ParAuth + auth.Email)
		rs.core.Error(err.Error(), zap.Stack("marshal_creation_failed"))
		return err
	}

	cell := types.Cell{
		Cell: register.Cell,
	}
	if err := rs.cellService.AssignCellToUser(cell, marshal.UserID); err != nil {
		rs.storage.DeleteUser(consts.ParUser + user.ID)
		rs.storage.DeleteAuth(consts.ParAuth + auth.Email)
		rs.storage.DeleteMarshal(consts.ParUser+user.ID, consts.ParMarshal+marshal.ID)
		rs.core.Error(err.Error(), zap.Stack("error_in_assigning_cell_to_user"))
		return err
	}

	return nil
}
