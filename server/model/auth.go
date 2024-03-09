package model

import (
	"errors"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
)

type Auth struct {
	Email         string `json:"email"`
	Password      string `json:"password,omitempty"`
	Token         string `json:"token"`
	Suspend       bool   `json:"suspend"`
	SuspendReason string `json:"suspend_reason"`
	UserID        string `json:"-"`
}

type AuthRepo struct {
	PK            string `dynamodbav:"PK"`
	SK            string `dynamodbav:"SK"`
	Password      string `dynamodbav:"Password"`
	Suspend       bool   `dynamodbav:"Suspend"`
	SuspendReason string `dynamodbav:"SuspendReason"`
	UserID        string `dynamodbav:"UserID"`
	EntityType    string `dynamodbav:"EntityType"`
}

func (a *Auth) ToRepo() AuthRepo {
	return AuthRepo{
		PK:            hashtag.Auth + a.Email,
		SK:            hashtag.Auth + a.Email,
		Password:      a.Password,
		Suspend:       a.Suspend,
		SuspendReason: a.SuspendReason,
		UserID:        a.UserID,
		EntityType:    entity.Auth,
	}
}

func (a *AuthRepo) ToAPI() Auth {
	return Auth{
		Email:         a.PK,
		Password:      a.Password,
		Suspend:       a.Suspend,
		SuspendReason: a.SuspendReason,
		UserID:        a.UserID,
	}
}

func (a *Auth) ValidateAuth() error {
	if err := validateEmail(a.Email); err != nil {
		return err
	}

	if !validatePassword(a.Password) {
		return errors.New("password not accepted")
	}

	return nil
}
