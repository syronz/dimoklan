package model

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"dimoklan/consts"
	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/internal/errors/errstatus"
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
	if err := validateEmail(r.Email); err != nil {
		return err
	}

	if err := validatePassword(r.Password); err != nil {
		return errors.New("password not accepted")
	}

	if !validateRegisterCell(r.Cell.ToString()) {
		return errors.New("cell is not valid")
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required; code: %w", errstatus.ErrNotAcceptable)
	}

	if email[0:2] != "a:" {
		return fmt.Errorf("email is follow pattern; code: %w", errstatus.ErrNotAcceptable)
	}

	regex := `^a:[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if match, _ := regexp.MatchString(regex, email); !match {
		return fmt.Errorf("email is not valid; code: %w", errstatus.ErrUnprocessableEntity)
	}

	emailSections := strings.Split(email, "@")
	if _, ok := consts.EmailProviders()[emailSections[1]]; !ok {
		return fmt.Errorf("email is not accepted; code: %w", errstatus.ErrForbidden)
	}

	return nil
}

func validatePassword(password string) error {
	// Check if the password length is at least 12 characters
	if password == "" {
		return fmt.Errorf("password is required; code: %w", errstatus.ErrNotAcceptable)
	}

	hasLength := false
	if len(password) >= consts.MinPasswordLength {
		hasLength = true
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
	if !(hasLength && hasLower && hasUpper && hasDigit) {
		return fmt.Errorf("password is not acceptable; code: %w", errstatus.ErrNotAcceptable)
	}

	return nil
}

func validateRegisterCell(cell string) bool {
	coords := strings.Split(cell, ":")
	if len(coords) != 3 {
		return false
	}
	if coords[0] != "c" {
		return false
	}
	num, err := strconv.Atoi(coords[1])
	if num == 0 || num > consts.MaxX || err != nil {
		return false
	}
	num, err = strconv.Atoi(coords[2])
	if num == 0 || num > consts.MaxX || err != nil {
		return false
	}

	return true
}
