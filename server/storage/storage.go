package storage

import (
	"dimoklan/types"
)

type Storage interface {
	Get(int) *types.User
}
