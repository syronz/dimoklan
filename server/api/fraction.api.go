package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"dimoklan/internal/config"
	"dimoklan/service"
	"dimoklan/model"
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
	fractions := strings.Split(c.QueryParam("coordinates"), ",")

	if len(fractions) == 0 {
		return c.JSON(http.StatusNotAcceptable, map[string]any{
			"error": "coordinates are required",
		})
	}

	fmt.Println(">>>> api", fractions)

	var err error
	var fractionData []model.Fraction
	if fractionData, err = s.fractionService.GetFractions(fractions); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, fractionData)
}
