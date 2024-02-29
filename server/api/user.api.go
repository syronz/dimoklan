package api

import (
	"dimoklan/internal/config"
	"dimoklan/service"
)

type UserAPI struct {
	core        config.Core
	userService *service.UserService
}

func NewUserAPI(core config.Core, userService *service.UserService) *UserAPI {
	return &UserAPI{
		core:        core,
		userService: userService,
	}
}
