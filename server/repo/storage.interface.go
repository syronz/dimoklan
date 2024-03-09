package repo

import (
	"context"

	"dimoklan/model"
)

type Storage interface {
	CreateUser(context.Context, model.User) error
	DeleteUser(context.Context, string) error
	GetUserByEmail(context.Context, string) (model.UserRepo, error)

	CreateRegister(context.Context, model.Register) error
	ConfirmRegister(context.Context, string) (model.RegisterRepo, error)

	CreateAuth(context.Context, model.Auth) error
	DeleteAuth(context.Context, string) error
	GetAuthByEmail(context.Context, string) (model.AuthRepo, error)

	CreateMarshal(context.Context, model.Marshal) error
	DeleteMarshal(context.Context, string, string) error

	CreateCell(context.Context, model.Cell) error
	GetCellByCoord(context.Context, int, int) (model.Cell, error)
	GetMapUsers(context.Context, model.Point, model.Point) (map[model.Point]int, error)

	GetFractions(context.Context, []string) ([]model.Fraction, error)
}
