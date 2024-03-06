package service

import (
	"fmt"

	"go.uber.org/zap"

	"dimoklan/internal/config"
	"dimoklan/repo"
	"dimoklan/model"
)

type FractionService struct {
	core    config.Core
	storage repo.Storage
}

func NewFractionService(core config.Core, storage repo.Storage) *FractionService {
	return &FractionService{
		core:    core,
		storage: storage,
	}
}

func (fs *FractionService) GetFractions(coordinates []string) ([]model.Fraction, error) {
	fmt.Println(">>>>>> service", coordinates)
	fractions, err := fs.storage.GetFractions(coordinates)
	if err != nil {
		fs.core.Error(err.Error(), zap.Stack("registration_failed"))
		return fractions, err
	}

	return fractions, nil
}
