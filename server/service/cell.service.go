package service

import (
	"fmt"
	"time"

	"dimoklan/internal/config"
	"dimoklan/repo"
	"dimoklan/types"

	"go.uber.org/zap"
)

type CellService struct {
	core    config.Core
	storage repo.Storage
}

func NewCellService(core config.Core, storage repo.Storage) *CellService {
	return &CellService{
		core:    core,
		storage: storage,
	}
}

func (s *CellService) GetCellByCoord(x, y int) (types.Cell, error) {
	cell, err := s.storage.GetCellByCoord(x, y)
	if err != nil {
		s.core.Error(err.Error(), zap.String("coordination", fmt.Sprintf("%v:%v", x, y)))
		return cell, err
	}

	return cell, nil
}

func (s *CellService) Create(cell types.Cell) (types.Cell, error) {
	cell.Fraction = cell.Cell.ToFraction()
	cell.UpdatedAt = time.Now().Unix()
	cell.Building = ""
	cell.Score = 10

	if err := s.storage.CreateCell(cell); err != nil {
		return cell, err
	}

	return cell, nil
}

func (s *CellService) AssignCellToUser(cell types.Cell, userID string) error {
	cell.Fraction = cell.Cell.ToFraction()
	cell.UpdatedAt = time.Now().Unix()
	cell.Score = 10
	cell.UserID = userID

	if err := s.storage.CreateCell(cell); err != nil {
		return err
	}

	return nil
}

/*
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
*/
