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
	"github.com/mailjet/mailjet-apiv3-go/v4"
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
	storage     *repo.Repo
	cellService *CellService
}

func NewRegisterService(core config.Core, storage *repo.Repo, cellService *CellService) *RegisterService {
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
	auth, err := rs.storage.GetAuthByEmail(ctx, register.Email)
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return register, err
	}

	if auth.Email != "" {
		return register, errors.New("email is not avaialble")
	}

	// TODO: this should be sent by email
	fmt.Println(">>>> actiation code: ", hex.EncodeToString(activationCode[:]))

	// Email
	mailjetClient := mailjet.NewMailjetClient(rs.core.GetMjApikeyPublic(), rs.core.GetMjApikeyPrivate())
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "noreply@erp14.click",
				Name:  "Dimoklan",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: register.Email[2:],
					Name:  register.Kingdom,
				},
			},
			Subject:  "Dimoklan: Activation Link",
			TextPart: "Welcome to Dimoklan! Click the link below to activate your account.",
			HTMLPart: `<h3>Activation Link: <a href="` + rs.core.GetAppURL() + `register?activation_code=` + hex.EncodeToString(activationCode[:]) + `">` + hex.EncodeToString(activationCode[:]) + `</a></h3><br />Dimoklan is a strategic game`,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("activation email not sent"))
		return register, err
	}

	fmt.Println("Activation Email Data: %+v", res)

	if err = rs.storage.CreateRegister(ctx, register); err != nil {
		rs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return register, err
	}

	register.Password = ""
	register.ActivationCode = ""
	return register, nil
}

func (rs *RegisterService) Confirm(ctx context.Context, activationCode string) error {
	activationCodeHashed := sha256.Sum256([]byte(activationCode))

	register, err := rs.storage.ConfirmRegister(ctx, hex.EncodeToString(activationCodeHashed[:]))
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("activation_failed"))
		return err
	}

	if register.Email == "" {
		return errors.New("activation is not valid")
	}

	// check if user already registered with same email
	savedAuth, err := rs.storage.GetAuthByEmail(ctx, register.Email)
	if err != nil {
		rs.core.Error(err.Error(), zap.Stack("activation_failed"))
		return err
	}

	if savedAuth.Email != "" {
		return errors.New("activation has already been completed")
	}

	randomID := rand.Intn(consts.MaxUserID)

	// create user
	user := model.User{
		ID:            fmt.Sprintf("%v%v", hashtag.User, randomID),
		Color:         strconv.FormatInt(int64(randomID), 16),
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

	if err := rs.storage.CreateUser(ctx, user); err != nil {
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

	if err := rs.storage.CreateAuth(ctx, auth); err != nil {
		rs.storage.DeleteUser(ctx, user.ID)
		rs.core.Error(err.Error(), zap.Stack("auth_creation_failed"))
		return err
	}

	// create a marshal for user
	marshal := model.Marshal{
		UserID:    user.ID,
		ID:        fmt.Sprintf("%v%v:1", hashtag.Marshal, randomID),
		Name:      sillyname.GenerateStupidName(),
		Cell:      register.Cell,
		Army:      newuser.Army,
		Star:      newuser.Star,
		Speed:     newuser.Speed,
		Attack:    newuser.Attack,
		Face:      "todo_to_be_added",
		CreatedAt: time.Now(),
	}

	if err := rs.storage.CreateMarshal(ctx, marshal); err != nil {
		rs.storage.DeleteUser(ctx, user.ID)
		rs.storage.DeleteAuth(ctx, auth.Email)
		rs.core.Error(err.Error(), zap.Stack("marshal_creation_failed"))
		return err
	}

	cell := model.Cell{
		Cell: register.Cell,
	}
	if err := rs.cellService.AssignCellToUser(ctx, cell, marshal.UserID); err != nil {
		rs.storage.DeleteUser(ctx, user.ID)
		rs.storage.DeleteAuth(ctx, auth.Email)
		rs.storage.DeleteMarshal(ctx, user.ID, marshal.ID, marshal.Cell.ToFraction())
		rs.core.Error(err.Error(), zap.Stack("error_in_assigning_cell_to_user"))
		return err
	}

	return nil
}
