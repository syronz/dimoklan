package types

import "errors"

type Auth struct {
	Email    string `json:"email" dynamodbav:"PK"`
	Password string `json:"password,omitempty" dynamodbav:"Password"`
	Token    string `json:"token" dynamodbav:"-"`
	SK       string `json:"-" dynamodbav:"SK"`
}

func (a *Auth) ValidateAuth() error {
	if !validateEmail(a.Email) {
		return errors.New("email is not valid")
	}

	if !validatePassword(a.Password) {
		return errors.New("password not accepted")
	}

	return nil
}
