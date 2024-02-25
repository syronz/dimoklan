package types

import (
	"errors"
	"regexp"
)

type Register struct {
	Email    string `json:"email"`
	Hash     string `json:"hash"`
	Kingdom  string `json:"kingdom"`
	Language string `json:"language"`
	Password string `json:"password"`
	TTL      int64  `json:"ttl"`
}

func (r *Register) ValidateCreate() error {
	if ok := isValidEmail(r.Email); !ok {
		return errors.New("email is not valid")
	}

	return nil
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}
