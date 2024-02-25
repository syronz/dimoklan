package basapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/service"
	"dimoklan/types"
)

type BasRegisterAPI struct {
	core config.Core
	registerService *service.RegisterService
}

func NewBasRegisterAPI(core config.Core, registerService *service.RegisterService) *BasRegisterAPI {
	return &BasRegisterAPI{
		core: core,
		registerService: registerService,
	}
}


func (s *BasRegisterAPI) CreateRegister(c echo.Context) error {
	var register types.Register

	if err := c.Bind(&register); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err,
		})
	}

	var err error
	if register, err = s.registerService.Create(register); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, register)
}
