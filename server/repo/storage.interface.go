package repo

import (
	"context"

	"dimoklan/model"
)

type Storage interface {
	CreateUser(context.Context, model.UserRepo) error
	DeleteUser(context.Context, string) error
	GetUserByEmail(context.Context, string) (model.UserRepo, error)

	CreateRegister(context.Context, model.RegisterRepo) error
	ConfirmRegister(context.Context, string) (model.RegisterRepo, error)

	CreateAuth(context.Context, model.AuthRepo) error
	DeleteAuth(context.Context, string) error
	GetAuthByEmail(context.Context, string) (model.Auth, error)

	CreateMarshal(context.Context, model.MarshalRepo) error
	DeleteMarshal(context.Context, string, string) error

	CreateCell(context.Context, model.CellRepo) error
	GetCellByCoord(context.Context, int, int) (model.Cell, error)
	GetMapUsers(context.Context, model.Point, model.Point) (map[model.Point]int, error)

	GetFractions(context.Context, []string) ([]model.Fraction, error)
}
