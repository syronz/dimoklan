package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"dimoklan/consts/ctxkey"
	"dimoklan/consts/entity"
	"dimoklan/internal/config"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/model"
	"dimoklan/repo"
)

type MarshalService struct {
	core    config.Core
	storage *repo.Repo
}

func NewMarshalService(core config.Core, storage *repo.Repo) *MarshalService {
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

	userID := ctx.Value(ctxkey.UserID)

	if move.UserID != userID {
		fs.core.Error("WARNING: HACK DTECTED", zap.Any("JWT user_id", userID), zap.String("payload user_id", move.UserID))
		return model.Marshal{}, fmt.Errorf("something went wrong; code: %w", errstatus.ErrNotAcceptable)
	}

	if err := fs.storage.UpdateEntityTypeMarshalMoving(ctx, move, entity.MarshalMoving); err != nil {
		fs.core.Error(err.Error(), zap.Stack("update entity type for marshal moving"))
		return model.Marshal{}, err
	}

	fmt.Printf(">>>>>>> 200: %+v\n", marshal)

	return marshal, nil
}
