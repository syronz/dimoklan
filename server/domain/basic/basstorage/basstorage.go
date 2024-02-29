package basstorage

import "dimoklan/types"

type BasStorage interface {
	CreateUser(types.User) error
	DeleteUser(string) error
	GetUserByEmail(string) (types.User, error)

	CreateRegister(types.Register) error
	ConfirmRegister(string) (types.Register, error)

	CreateAuth(types.Auth) error
	DeleteAuth(string) error
	GetAuthByEmail(string) (types.Auth, error)

	CreateMarshal(types.Marshal) error
	DeleteMarshal(string, string) error
}
