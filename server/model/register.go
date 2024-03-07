package model

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"dimoklan/consts"
	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/model/localtype"
)

type Register struct {
	Email          string         `json:"email"`
	Kingdom        string         `json:"kingdom"`
	Language       string         `json:"language"`
	Password       string         `json:"password,omitempty"`
	ActivationCode string         `json:"activation_code,omitempty"`
	Cell           localtype.CELL `json:"cell,omitempty"`
}

type RegisterRepo struct {
	PK             string
	SK             string
	EntityType     string
	Email          string
	Cell           localtype.CELL
	Kingdom        string
	Language       string
	Password       string
	ActivationCode string
	TTL            int64
}

func (r *Register) ToRepo() RegisterRepo {
	return RegisterRepo{
		PK:         hashtag.Register + r.ActivationCode,
		SK:         hashtag.Register + r.ActivationCode,
		EntityType: entity.Register,
		Email:      r.Email,
		TTL:        time.Now().Add(24 * time.Hour).Unix(),
		Kingdom:    r.Kingdom,
		Language:   r.Language,
		Password:   r.Password,
		Cell:       r.Cell,
	}
}

func (r *RegisterRepo) ToAPI() Register {
	return Register{
		Email:          r.Email,
		Kingdom:        r.Kingdom,
		Language:       r.Language,
		Password:       r.Password,
		ActivationCode: r.ActivationCode,
		Cell:           r.Cell,
	}
}

func (r *Register) ValidateRegister() error {
	if !validateEmail(r.Email) {
		return errors.New("email is not valid")
	}

	if !validatePassword(r.Password) {
		return errors.New("password not accepted")
	}

	if !validateRegisterCell(r.Cell.ToString()) {
		return errors.New("cell is not valid")
	}

	return nil
}

func validateEmail(email string) bool {
	if email == "" {
		return false
	}

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

func validateRegisterCell(cell string) bool {
	coords := strings.Split(cell, ":")
	if len(coords) != 2 {
		return false
	}
	num, err := strconv.Atoi(coords[0])
	if num == 0 || num > consts.MaxX || err != nil {
		return false
	}
	num, err = strconv.Atoi(coords[1])
	if num == 0 || num > consts.MaxX || err != nil {
		return false
	}

	return true
}
