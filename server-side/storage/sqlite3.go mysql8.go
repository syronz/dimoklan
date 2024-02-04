package storage

import (
	"dimoklan/types"
)

type MySQL struct{}

func (s *MySQL) Get(id int) *types.User {
	return &types.User{
		Bit:  1,
		Name: "Adrian",
	}
}
