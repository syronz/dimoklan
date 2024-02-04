package basstorage

import "dimoklan/types"

type BasStorage interface {
	CreateUser(types.User) error
	GetUserByColor(string) (types.User, error)
}
