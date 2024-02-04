package storage

import (
	"dimoklan/types"
)

type MySQL struct{}

func (s *MySQL) Get(id int) *types.User {
	return &types.User{
		ID:  1,
		Name: "Adrian",
	}
}
