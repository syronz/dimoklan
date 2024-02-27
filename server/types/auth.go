package types

import "errors"

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Auth) validateAuth() error {
	if !validateEmail(a.Email) {
		return errors.New("email is not valid")
	}

	if !validatePassword(a.Password) {
		return errors.New("password not accepted")
	}

	return nil
}
