package types

import "errors"

type Auth struct {
	Email         string `json:"email" dynamodbav:"PK"`
	Password      string `json:"password,omitempty" dynamodbav:"Password"`
	Token         string `json:"token" dynamodbav:"-"`
	Suspend       bool   `json:"suspend" dynamodbav:"Suspend"`
	SuspendReason string `json:"suspend_reason" dynamodbav:"SuspendReason"`
	UserID        string `json:"-" dynamodbav:"UserID"`
	SK            string `json:"-" dynamodbav:"SK"`
	EntityType    string `json:"-" dynamodbav:"EntityType"`
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
