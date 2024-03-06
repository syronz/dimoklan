package service

import (
	"fmt"
	"time"

	"dimoklan/internal/config"
	"dimoklan/repo"
	"dimoklan/model"

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

func (s *CellService) GetCellByCoord(x, y int) (model.Cell, error) {
	cell, err := s.storage.GetCellByCoord(x, y)
	if err != nil {
		s.core.Error(err.Error(), zap.String("coordination", fmt.Sprintf("%v:%v", x, y)))
		return cell, err
	}

	return cell, nil
}

func (s *CellService) Create(cell model.Cell) (model.Cell, error) {
	cell.Fraction = cell.Cell.ToFraction()
	cell.UpdatedAt = time.Now().Unix()
	cell.Building = ""
	cell.Score = 10

	if err := s.storage.CreateCell(cell); err != nil {
		return cell, err
	}

	return cell, nil
}

func (s *CellService) AssignCellToUser(cell model.Cell, userID string) error {
	cell.Fraction = cell.Cell.ToFraction()
	cell.UpdatedAt = time.Now().Unix()
	cell.Score = 10
	cell.UserID = userID

	if err := s.storage.CreateCell(cell); err != nil {
		return err
	}

	return nil
}
