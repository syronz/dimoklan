package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/service"
)

type MarshalAPI struct {
	core           config.Core
	marshalService *service.MarshalService
}

func NewMarshalAPI(core config.Core, marshalService *service.MarshalService) *MarshalAPI {
	return &MarshalAPI{
		core:           core,
		marshalService: marshalService,
	}
}

func (s *MarshalAPI) GetMarshal(c echo.Context) error {
	id := c.Param("id")

	marshal, err := s.marshalService.GetMarshal(c.Request().Context(), id)
	if err != nil {
		return c.JSON(status(err), map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, marshal)
}

func (s *MarshalAPI) MoveMarshal(c echo.Context) error {
	var move model.Move

	if err := c.Bind(&move); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]any{
			"error": err,
		})
	}

	move.MarshalID = c.Param("id")

	marshal, err := s.marshalService.MoveMarshal(c.Request().Context(), move)
	if err != nil {
		return c.JSON(status(err), map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, marshal)
}
