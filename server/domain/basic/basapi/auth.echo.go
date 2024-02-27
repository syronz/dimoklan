package basapi

import (
	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/service"
)

type BasAuthAPI struct {
	core        config.Core
	userService *service.UserService
}

func NewBasAuthAPI(core config.Core, userService *service.UserService) *BasAuthAPI {
	return &BasAuthAPI{
		core:        core,
		userService: userService,
	}
}

func (s *BasAuthAPI) CreateAuth(c echo.Context) error {
	// var auth types.Auth

	// if err := c.Bind(&auth); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"error": err,
	// 	})
	// }

	// var err error
	// if auth, err = s.userService.Create(auth); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]any{
	// 		"error": err.Error(),
	// 	})
	// }

	// return c.JSON(http.StatusOK, auth)
	return nil
}
