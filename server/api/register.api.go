package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/service"
)

type RegisterAPI struct {
	core            config.Core
	registerService *service.RegisterService
}

func NewRegisterAPI(core config.Core, registerService *service.RegisterService) *RegisterAPI {
	return &RegisterAPI{
		core:            core,
		registerService: registerService,
	}
}

func (br *RegisterAPI) Create(c echo.Context) error {
	var register model.Register

	if err := c.Bind(&register); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err,
		})
	}

	var err error
	if register, err = br.registerService.Create(c.Request().Context(), register); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, register)
}

func printMessage(head, content string) string {
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Email confirmation</title>
		</head>
		<body>
			<h1>%v</h1>
			<p>%v</p>
		</body>
		</html>
		`, head, content)
}

func (br *RegisterAPI) Confirm(c echo.Context) error {
	hashCode := c.QueryParam("activation_code")

	if hashCode == "" {
		return c.HTML(http.StatusConflict, printMessage("Link is not valid", ""))
	}

	var err error
	if err = br.registerService.Confirm(c.Request().Context(), hashCode); err != nil {
		return c.HTML(http.StatusOK, printMessage("Confirmation Failed", err.Error()))
	}

	htmlContent := printMessage("Confirmation Successful", fmt.Sprintf(`Please go to <a href="%v">Log in</a>`, br.core.GetLoginPage()))
	return c.HTML(http.StatusOK, htmlContent)
}
