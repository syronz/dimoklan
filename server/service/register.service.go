package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.uber.org/zap"

	"dimoklan/consts"
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/types"
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
	if err := register.ValidateCreate(); err != nil {
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
	password := fmt.Sprintf("%v%v%v", consts.HashSalt, register.Password, rs.core.GetSalt())
	hashedPassword := sha256.Sum256([]byte(password))
	register.Password = hex.EncodeToString(hashedPassword[:])

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
		rs.core.Error(err.Error(), zap.Stack("confirmation_failed"))
		return err
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
