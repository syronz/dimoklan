package storage

import (
	"dimoklan/types"
)

type MemoryStorage struct{}

func NewMemroryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		Bit:  1,
		Name: "Adrian",
	}
}