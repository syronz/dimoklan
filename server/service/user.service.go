package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"dimoklan/consts"
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/types"

	"github.com/google/uuid"
	"go.uber.org/zap"
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

func (s *UserService) GetUserByColor(color string) (types.User, error) {
	user, err := s.storage.GetUserByColor(color)
	if err != nil {
		s.core.Error(err.Error(), zap.String("color", color))
		return user, err
	}

	return user, nil
}

func (s *UserService) Create(user types.User) (types.User, error) {
	color, err := s.pickColor(); 
	if err != nil {
		s.core.Error(err.Error(), zap.Stack("color_conflict_in_creating_user"))
		return user, err
	}

	user.Code = uuid.New().String()
	user.Color = color
	user.Status = consts.Active
	user.Reason = "new user"

	if err := s.storage.CreateUser(user); err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserService) pickColor() (string, error) {
	var color string
	for i := 0; i <= consts.MaxRetryForPickColor; i++ {
		color = generateRandomColor()
		existedUser, err := s.GetUserByColor(color)
		if err != nil {
			return "", err
		}

		if existedUser.ID == 0 {
			break
		}

		if i == consts.MaxRetryForPickColor {
			err = errors.New("failed to generate non existed color")
			return "", err
		}
	}

	return color, nil
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
