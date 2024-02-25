package service

import (
	"crypto/sha256"
	"encoding/hex"
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

func (s *RegisterService) Create(register types.Register) (types.Register, error) {
	if err := register.ValidateCreate(); err != nil {
		return types.Register{}, err
	}


	// Get the current time
	currentTime := time.Now().String() + consts.HashSalt

	// Convert the current time to a byte slice
	currentTimeBytes := []byte(currentTime)

	// Calculate the SHA-256 hash of the byte slice
	hash1 := sha256.Sum256(currentTimeBytes)
	hash2 := sha256.Sum256(hash1[:])

	register.Hash = hex.EncodeToString(hash2[:])
	// delete after 24 hours
	register.TTL = time.Now().Add(24 * time.Hour).Unix() 
	register.Language = consts.LanguageEn

	if err := s.storage.CreateRegister(register); err != nil {
		s.core.Error(err.Error(), zap.Stack("registration_failed"))
		return register, err
	}

	return register, nil
}
