package service

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"dimoklan/cache"
	"dimoklan/consts/ctxkey"
	"dimoklan/consts/gp"
	"dimoklan/internal/config"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/model"
	"dimoklan/repo"
)

type MarshalService struct {
	core    config.Core
	cache   *cache.Cache
	storage *repo.Repo
}

func NewMarshalService(core config.Core, cache *cache.Cache, storage *repo.Repo) *MarshalService {
	return &MarshalService{
		core:    core,
		cache:   cache,
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

	if marshal.UserID != userID {
		fs.core.Error("WARNING: HACK DTECTED", zap.Any("JWT user_id", userID), zap.String("payload user_id", marshal.UserID))
		return model.Marshal{}, fmt.Errorf("something went wrong; code: %w", errstatus.ErrNotAcceptable)
	}

	// Check if in-progress move exist or not.
	ongoingMM, err := fs.cache.GetMarshalMove(ctx, move.MarshalID)
	if err != nil {
		fs.core.Error(err.Error(), zap.Stack("getting_marshal_move_failed_in_move_marshal"))
		return model.Marshal{}, fmt.Errorf("something went wrong; code: %w", errstatus.ErrNotAcceptable)
	}

	if ongoingMM.ArriveAt < time.Now().UnixMilli() {
		if err := fs.cache.DeleteMarshalMove(
			ctx,
			move.MarshalID,
			ongoingMM.Source,
			ongoingMM.Destination,
		); err != nil {
			fs.core.Error(err.Error(), zap.Stack("delete_marshal_move_failed_in_move_marshal"))
			return model.Marshal{}, fmt.Errorf("something went wrong; code: %w", errstatus.ErrNotAcceptable)

		}
	}

	if err := fs.SaveMove(ctx, move, marshal); err != nil {
		return model.Marshal{}, err
	}

	return marshal, nil
}

func (fs *MarshalService) SaveMove(ctx context.Context, move model.Move, marshal model.Marshal) error {
	marshalMove := model.MarshalMove{
		MarshalID:   marshal.ID,
		UserID:      marshal.UserID,
		Name:        marshal.Name,
		Star:        marshal.Star,
		Speed:       marshal.Speed,
		Face:        marshal.Face,
		Source:      marshal.Cell.ToString(),
		Destination: move.Cell.ToString(),
		DepartureAt: time.Now().UnixMilli(),
		ArriveAt:    time.Now().UnixMilli(),
	}

	// source
	marshalMove.Directrion = gp.Source
	if err := fs.cache.AddMarshalMoveToFraction(ctx, marshal.Cell.ToFraction(), marshalMove); err != nil {
		fs.core.Error(err.Error(), zap.Stack("add_move_marshal_source_failed"))
		return err
	}

	if move.Cell.ToFraction() != marshal.Cell.ToFraction() {
		// destination recorded if it is belong to another fraction
		marshalMove.Directrion = gp.Destination
		if err := fs.cache.AddMarshalMoveToFraction(ctx, move.Cell.ToFraction(), marshalMove); err != nil {
			fs.core.Error(err.Error(), zap.Stack("add_move_marshal_source_failed"))
			return err
		}
	}

	if err := fs.cache.AddMarshalMove(ctx, marshalMove); err != nil {
		fs.core.Error(err.Error(), zap.Stack("add_marshal_move_failed"))
		return err
	}

	return nil
}
