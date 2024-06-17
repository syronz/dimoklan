package model

import (
	"time"

	"dimoklan/consts/entity"
)

type User struct {
	ID            string    `json:"id"`
	Color         string    `json:"color"`
	Farr          int       `json:"farr"`
	Gold          int       `json:"gold"`
	Email         string    `json:"email"`
	Kingdom       string    `json:"kingdom"`
	Password      string    `json:"password,omitempty"`
	Language      string    `json:"language"`
	Suspend       bool      `json:"suspend"`
	SuspendReason string    `json:"suspend_reason"`
	Freeze        bool      `json:"freeze"`
	FreezeReason  string    `json:"freeze_reason"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserRepo struct {
	PK            string
	SK            string
	Color         string
	Farr          int
	Gold          int
	Email         string
	Kingdom       string
	Password      string
	Language      string
	Suspend       bool
	SuspendReason string
	Freeze        bool
	FreezeReason  string
	CreatedAt     int64
	UpdatedAt     int64
	EntityType    string
}

func (u *User) ToRepo() UserRepo {
	return UserRepo{
		PK:            u.ID,
		SK:            u.ID,
		Color:         u.Color,
		Farr:          u.Farr,
		Gold:          u.Gold,
		Email:         u.Email,
		Kingdom:       u.Kingdom,
		Password:      u.Password,
		Language:      u.Language,
		Suspend:       u.Suspend,
		SuspendReason: u.SuspendReason,
		Freeze:        u.Freeze,
		FreezeReason:  u.FreezeReason,
		CreatedAt:     u.CreatedAt.Unix(),
		UpdatedAt:     u.UpdatedAt.Unix(),
		EntityType:    entity.User,
	}
}

func (u *UserRepo) ToAPI() User {
	return User{
		ID:            u.PK,
		Color:         u.Color,
		Farr:          u.Farr,
		Gold:          u.Gold,
		Email:         u.Email,
		Kingdom:       u.Kingdom,
		Password:      "",
		Language:      u.Language,
		Suspend:       u.Suspend,
		SuspendReason: u.SuspendReason,
		Freeze:        u.Freeze,
		FreezeReason:  u.FreezeReason,
		CreatedAt:     time.Unix(u.CreatedAt, 0),
		UpdatedAt:     time.Unix(u.UpdatedAt, 0),
	}
}

func validateUser(u *User) bool { return true }
