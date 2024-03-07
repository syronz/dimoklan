package service

import (
	"context"
	"errors"
	"time"

	"dimoklan/consts"
	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/repo"
	"dimoklan/util"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type AuthService struct {
	core    config.Core
	storage repo.Storage
}

func NewAuthService(core config.Core, storage repo.Storage) *AuthService {
	return &AuthService{
		core:    core,
		storage: storage,
	}
}

func (as *AuthService) Login(ctx context.Context, auth model.Auth) (model.Auth, error) {
	if err := auth.ValidateAuth(); err != nil {
		return model.Auth{}, err
	}

	savedAuth, err := as.storage.GetAuthByEmail(ctx, auth.Email)
	if err != nil {
		as.core.Error(err.Error(), zap.Stack("get_auth_by_email_failed"))
		return model.Auth{}, err
	}

	if savedAuth.Email == "" {
		return model.Auth{}, errors.New("email or password is wrong")
	}

	if savedAuth.Password != util.HashPassword(auth.Password, consts.HashSalt, as.core.GetSalt()) {
		return model.Auth{}, errors.New("email or password is wrong")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": auth.UserID,
		"nbf":     time.Now().Unix(),
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(consts.HashSalt + as.core.GetSalt()))
	if err != nil {
		as.core.Error(err.Error(), zap.Stack("token_generation_failed"))
		return model.Auth{}, err
	}

	auth.Password = ""
	auth.Token = tokenString
	return auth, nil
}
