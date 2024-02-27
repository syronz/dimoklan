package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.uber.org/zap"

	"dimoklan/consts"
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/types"
	"dimoklan/util"
)

type RegisterService struct {
	core    config.Core
	storage basstorage.BasStorage
}

func NewRegisterService(core config.Core, storage basstorage.BasStorage) *RegisterService {
	return &RegisterService{
		core:    core,
		storage: storage,
	}
}

func (rs *RegisterService) Create(register types.Register) (types.Register, error) {
	if err := register.ValidateRegister(); err != nil {
		return types.Register{}, err
	}

	// Get the current time
	currentTime := time.Now().String() + consts.HashSalt

	// Convert the current time to a byte slice
	currentTimeBytes := []byte(currentTime)

	// Calculate the SHA-256 hash of the byte slice
	activationCode := sha256.Sum256(currentTimeBytes)
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

	fmt.Println(">>>> actiation code: ", hex.EncodeToString(activationCode[:]))

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

	user := types.User{
		ID:            id,
		Color:         strconv.FormatInt(int64(id), 16),
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

	return nil
}
