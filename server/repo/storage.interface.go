package repo

import (
	"context"

	"dimoklan/model"
)

type Storage interface {
	CreateUser(model.UserRepo) error
	DeleteUser(string) error
	GetUserByEmail(string) (model.UserRepo, error)

	CreateRegister(context.Context, model.RegisterRepo) error
	ConfirmRegister(context.Context, string) (model.RegisterRepo, error)

	CreateAuth(model.Auth) error
	DeleteAuth(string) error
	GetAuthByEmail(string) (model.Auth, error)

	CreateMarshal(model.Marshal) error
	DeleteMarshal(string, string) error

	CreateCell(model.Cell) error
	GetCellByCoord(int, int) (model.Cell, error)
	GetMapUsers(model.Point, model.Point) (map[model.Point]int, error)

	GetFractions([]string) ([]model.Fraction, error)
}
