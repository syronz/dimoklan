package service

import (
	"fmt"
	"time"

	"dimoklan/domain/map/mapstorage"
	"dimoklan/internal/config"
	"dimoklan/types"
	"dimoklan/util"

	"go.uber.org/zap"
)

type CellService struct {
	core        config.Core
	mapStorage  mapstorage.MapStorage
	userService *UserService
}

func NewCellService(
	core config.Core,
	storage mapstorage.MapStorage,
	userService *UserService,
) *CellService {
	return &CellService{
		core:        core,
		mapStorage:  storage,
		userService: userService,
	}
}

func (s *CellService) GetCellByCoord(x, y int) (types.Cell, error) {
	cell, err := s.mapStorage.GetCellByCoord(x, y)
	if err != nil {
		s.core.Error(err.Error(), zap.String("coordination", fmt.Sprintf("%v:%v", x, y)))
		return cell, err
	}

	return cell, nil
}

func toFraction(x,y int) string {
	x =          util.CeilInt(float64(x) / 10)
	y =          util.CeilInt(float64(y) / 10)

	return fmt.Sprintf("%d:%d", x, y)
}

func (s *CellService) Create(cell types.Cell) (types.Cell, error) {
	cell.Fraction = toFraction(cell.X, cell.Y)
	cell.Cell = fmt.Sprintf("%03d:%03d", cell.X, cell.Y)
	cell.LastUpdate = time.Now().Unix()
	cell.Building = ""
	cell.Score = 10

	if err := s.mapStorage.CreateCell(cell); err != nil {
		return cell, err
	}

	return cell, nil
}

func (s *CellService) GetMap(start, stop types.Point) ([][]string, error) {
	mapUsers, err := s.mapStorage.GetMapUsers(start, stop)
	if err != nil {
		s.core.Error(err.Error(), zap.String("area", fmt.Sprintf("%v:%v", start, stop)))
		return nil, err
	}

	mapColors, err := s.userService.GetAllColors()
	if err != nil {
		s.core.Error(err.Error(), zap.String("area", fmt.Sprintf("%v:%v", start, stop)))
		return nil, err
	}

	pixels := make([][]string, stop.Y-start.Y)
	for i := range pixels {
		pixels[i] = make([]string, stop.X-start.X)
	}

	for x := start.X; x < stop.X; x++ {
		for y := start.Y; y < stop.Y; y++ {
			userID, ok := mapUsers[types.Point{X: x, Y: y}]
			if !ok {
				continue
			}

			pixels[x][y] = mapColors[userID]
		}
	}

	return pixels, nil
}
