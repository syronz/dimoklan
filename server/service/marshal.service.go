package service

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"dimoklan/cache"
	"dimoklan/consts/ctxkey"
	"dimoklan/internal/config"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/model"
	"dimoklan/repo"
)

const (
	source      = "source"
	destination = "destination"
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

	if err := fs.SaveMove(ctx, move, marshal); err != nil {
		return model.Marshal{}, err
	}

	return marshal, nil
}

func (fs *MarshalService) SaveMove(ctx context.Context, move model.Move, marshal model.Marshal) error {
	moveMarshal := model.MoveMarshal{
		MarshalID:   marshal.ID,
		UserID:      marshal.UserID,
		Name:        marshal.Name,
		Star:        marshal.Star,
		Speed:       marshal.Speed,
		Face:        marshal.Face,
		Source:      marshal.Cell.ToString(),
		Destination: move.Cell.ToString(),
		DepartureAt: time.Now(),
		ArriveAt:    time.Now().Add(2 * time.Minute),
	}

	// source
	moveMarshal.Directrion = source
	if err := fs.cache.AddMarshalMoveToFraction(ctx, marshal.Cell.ToFraction(), moveMarshal); err != nil {
		fs.core.Error(err.Error(), zap.Stack("add_move_marshal_source_failed"))
		return err
	}

	if move.Cell.ToFraction() != marshal.Cell.ToFraction() {
		// destination recorded if it is belong to another fraction
		moveMarshal.Directrion = destination
		if err := fs.cache.AddMarshalMoveToFraction(ctx, move.Cell.ToFraction(), moveMarshal); err != nil {
			fs.core.Error(err.Error(), zap.Stack("add_move_marshal_source_failed"))
			return err
		}
	}

	if err := fs.cache.AddMarshalMove(ctx, moveMarshal); err != nil {
		fs.core.Error(err.Error(), zap.Stack("add_marshal_move_failed"))
		return err
	}

	return nil
}
