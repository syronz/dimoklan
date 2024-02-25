package basapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/service"
	"dimoklan/types"
)

type BasUserAPI struct {
	core config.Core
	userService *service.UserService
}

func NewBasUserAPI(core config.Core, userService *service.UserService) *BasUserAPI {
	return &BasUserAPI{
		core: core,
		userService: userService,
	}
}


func (s *BasUserAPI) CreateUser(c echo.Context) error {
	var user types.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err,
		})
	}

	var err error
	if user, err = s.userService.Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}
