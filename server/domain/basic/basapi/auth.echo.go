package basapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/service"
	"dimoklan/types"
)

type BasAuthAPI struct {
	core        config.Core
	authService *service.AuthService
}

func NewBasAuthAPI(core config.Core, authService *service.AuthService) *BasAuthAPI {
	return &BasAuthAPI{
		core:        core,
		authService: authService,
	}
}

func (s *BasAuthAPI) Login(c echo.Context) error {
	var auth types.Auth

	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err,
		})
	}

	var err error
	if auth, err = s.authService.Login(auth); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, auth)
}
