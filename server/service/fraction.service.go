package service

import (
	"context"

	"go.uber.org/zap"

	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/repo"
)

type FractionService struct {
	core    config.Core
	storage *repo.Repo
}

func NewFractionService(core config.Core, storage *repo.Repo) *FractionService {
	return &FractionService{
		core:    core,
		storage: storage,
	}
}

func (fs *FractionService) GetFractions(ctx context.Context, coordinates []string) ([]model.Fraction, error) {
	fractions, err := fs.storage.GetFractions(ctx, coordinates)
	if err != nil {
		fs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return nil, err
	}

	return fractions, nil
}
