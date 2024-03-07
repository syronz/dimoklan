package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/service"
)

type FractionAPI struct {
	core            config.Core
	fractionService *service.FractionService
}

func NewFractionAPI(core config.Core, fractionService *service.FractionService) *FractionAPI {
	return &FractionAPI{
		core:            core,
		fractionService: fractionService,
	}
}

func (s *FractionAPI) GetFraction(c echo.Context) error {
	coordinates := strings.Split(c.QueryParam("coordinates"), ",")

	if len(coordinates) == 0 {
		return c.JSON(http.StatusNotAcceptable, map[string]any{
			"error": "coordinates are required",
		})
	}

	fractions, err := s.fractionService.GetFractions(c.Request().Context(), coordinates)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, fractions)
}
