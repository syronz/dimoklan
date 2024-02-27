package types

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"dimoklan/consts"
)

type Register struct {
	Email          string `json:"email"`
	ActivationCode string `json:"activation_code"`
	Kingdom        string `json:"kingdom"`
	Language       string `json:"language"`
	Password       string `json:"password"`
	TTL            int64  `json:"ttl"`
}

func (r *Register) ValidateRegister() error {
	if !validateEmail(r.Email) {
		return errors.New("email is not valid")
	}

	if !validatePassword(r.Password) {
		return errors.New("password not accepted")
	}

	return nil
}

func validateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if match, _ := regexp.MatchString(regex, email); !match {
		return false
	}

	emailSections := strings.Split(email, "@")
	if _, ok := consts.EmailProviders()[emailSections[1]]; !ok {
		return false
	}

	return true
}

func validatePassword(password string) bool {
	// Check if the password length is at least 12 characters
	if len(password) < 12 {
		return false
	}

	// Check if the password contains at least one lowercase letter
	hasLower := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
			break
		}
	}

	// Check if the password contains at least one uppercase letter
	hasUpper := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
			break
		}
	}

	// Check if the password contains at least one digit
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}

	// Return true if all criteria are met
	return hasLower && hasUpper && hasDigit
}
