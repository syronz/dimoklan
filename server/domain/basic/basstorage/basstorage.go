package basstorage

import "dimoklan/types"

type BasStorage interface {
	CreateUser(types.User) error
	GetUserByEmail(string) (types.User, error)

	CreateRegister(types.Register) error
	ConfirmRegister(string) (types.Register, error)
}
