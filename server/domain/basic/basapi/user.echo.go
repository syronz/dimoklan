package basapi

import (
	"dimoklan/internal/config"
	"dimoklan/service"
)

type BasUserAPI struct {
	core        config.Core
	userService *service.UserService
}

func NewBasUserAPI(core config.Core, userService *service.UserService) *BasUserAPI {
	return &BasUserAPI{
		core:        core,
		userService: userService,
	}
}
