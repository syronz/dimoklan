package service

import (
	"dimoklan/internal/config"
	"dimoklan/repo"
)

type UserService struct {
	core    config.Core
	storage *repo.Repo
}

func NewUserService(core config.Core, storage *repo.Repo) *UserService {
	return &UserService{
		core:    core,
		storage: storage,
	}
}
