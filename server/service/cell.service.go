package service

import (
	"dimoklan/domain/map/mapstorage"
	"dimoklan/internal/config"
	"dimoklan/types"
	"fmt"

	"go.uber.org/zap"
)

type CellService struct {
	core    config.Core
	storage mapstorage.MapStorage
}

func NewCellService(core config.Core, storage mapstorage.MapStorage) *CellService {
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
	cell.Building = ""
	cell.Score = 10

	if err := s.storage.CreateCell(cell); err != nil {
		return cell, err
	}

	return cell, nil
}
