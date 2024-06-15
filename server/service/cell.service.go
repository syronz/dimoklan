package service

import (
	"context"
	"fmt"
	"time"

	"dimoklan/consts/newuser"
	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/repo"

	"go.uber.org/zap"
)

type CellService struct {
	core    config.Core
	storage *repo.Repo
}

func NewCellService(core config.Core, storage *repo.Repo) *CellService {
	return &CellService{
		core:    core,
		storage: storage,
	}
}

func (s *CellService) GetCellByCoord(ctx context.Context, x, y int) (model.Cell, error) {
	cell, err := s.storage.GetCellByCoord(ctx, x, y)
	if err != nil {
		s.core.Error(err.Error(), zap.String("coordination", fmt.Sprintf("%v:%v", x, y)))
		return cell, err
	}

	return cell, nil
}

func (s *CellService) Create(ctx context.Context, cell model.Cell) (model.Cell, error) {
	cell.Fraction = cell.Cell.ToFraction()
	cell.UpdatedAt = time.Now()
	cell.Building = ""
	cell.Score = newuser.Score

	if err := s.storage.CreateCell(ctx, cell); err != nil {
		return cell, err
	}

	return cell, nil
}

func (s *CellService) AssignCellToUser(ctx context.Context, cell model.Cell, userID string) error {
	cell.Fraction = cell.Cell.ToFraction()
	cell.UpdatedAt = time.Now()
	cell.Score = newuser.Score
	cell.UserID = userID

	if err := s.storage.CreateCell(ctx, cell); err != nil {
		return err
	}

	return nil
}
