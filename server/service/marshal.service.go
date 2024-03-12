package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/repo"
)

type MarshalService struct {
	core    config.Core
	storage repo.Storage
}

func NewMarshalService(core config.Core, storage repo.Storage) *MarshalService {
	return &MarshalService{
		core:    core,
		storage: storage,
	}
}

func (fs *MarshalService) GetMarshal(ctx context.Context, id string) (model.Marshal, error) {

	marshal, err := fs.storage.GetMarshal(ctx, id)
	if err != nil {
		fs.core.Error(err.Error(), zap.Stack("getting_marshal_failed"))
		return model.Marshal{}, err
	}

	return marshal, nil
}

func (fs *MarshalService) MoveMarshal(ctx context.Context, move model.Move) (model.Marshal, error) {
	if err := move.Validate(); err != nil {
		return model.Marshal{}, err
	}

	marshal, err := fs.storage.GetMarshal(ctx, move.MarshalID)
	if err != nil {
		fs.core.Error(err.Error(), zap.Stack("getting_marshal_failed_in_move"))
		return model.Marshal{}, err
	}

	// userID := ctx.Value(consts.UserID)

	fmt.Printf(">>>>>>> 199: %+v\n", ctx)

	return marshal, nil
}
