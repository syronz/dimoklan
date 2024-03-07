package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/service"
)

type AuthAPI struct {
	core        config.Core
	authService *service.AuthService
}

func NewAuthAPI(core config.Core, authService *service.AuthService) *AuthAPI {
	return &AuthAPI{
		core:        core,
		authService: authService,
	}
}

func (s *AuthAPI) Login(c echo.Context) error {
	var auth model.Auth

	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err,
		})
	}

	var err error
	if auth, err = s.authService.Login(c.Request().Context(), auth); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, auth)
}
