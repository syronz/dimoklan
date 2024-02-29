package service

import (
	"crypto/rand"
	"encoding/hex"

	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
)

type UserService struct {
	core    config.Core
	storage basstorage.BasStorage
}

func NewUserService(core config.Core, storage basstorage.BasStorage) *UserService {
	return &UserService{
		core:    core,
		storage: storage,
	}
}

func generateRandomColor() string {
	// Determine the number of bytes needed for a 6-character hex
	numBytes := 3

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// TODO: Handle the error (in a real application, you would handle errors appropriately)
		return "000000" // Default color in case of error
	}

	// Convert to hexadecimal string
	hexString := hex.EncodeToString(randomBytes)

	// Add a '#' prefix
	return hexString
}
