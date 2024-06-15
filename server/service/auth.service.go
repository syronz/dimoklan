package service

import (
	"context"
	"fmt"
	"time"

	"dimoklan/consts"
	"dimoklan/internal/config"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/model"
	"dimoklan/repo"
	"dimoklan/util"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type AuthService struct {
	core    config.Core
	storage *repo.Repo
}

func NewAuthService(core config.Core, storage *repo.Repo) *AuthService {
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
		return model.Auth{}, fmt.Errorf("email or password is wrong; code: %w", errstatus.ErrForbidden)
	}

	if savedAuth.Password != util.HashPassword(auth.Password, consts.HashSalt, as.core.GetSalt()) {
		return model.Auth{}, fmt.Errorf("email or password is wrong; code: %w", errstatus.ErrForbidden)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		consts.UserID: savedAuth.UserID,
		"nbf":         time.Now().Unix(),
		"exp":         time.Now().Add(70000 * 24 * time.Hour).Unix(),
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
