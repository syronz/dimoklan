package repo

import (
	"context"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/model"
)

func (r *Repo) CreateRegister(ctx context.Context, register model.Register) error {
	return r.putUniqueItem(ctx, entity.Register, register.ToRepo())
}

func (r *Repo) ConfirmRegister(ctx context.Context, activationCode string) (model.Register, error) {
	registerRepo := model.RegisterRepo{}
	if err := r.getItem(ctx, entity.Register, &registerRepo, hashtag.Register+activationCode); err != nil {
		return model.Register{}, err
	}

	return registerRepo.ToAPI(), nil
}
