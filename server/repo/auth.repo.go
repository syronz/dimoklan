package repo

import (
	"context"

	"dimoklan/consts/entity"
	"dimoklan/model"
)

func (r *Repo) CreateAuth(ctx context.Context, auth model.Auth) error {
	return r.putUniqueItem(ctx, entity.Auth, auth.ToRepo())
}

func (r *Repo) DeleteAuth(ctx context.Context, authID string) error {
	pk := authID
	return r.deleteItem(ctx, entity.Auth, pk)
}

func (r *Repo) GetAuthByEmail(ctx context.Context, email string) (model.Auth, error) {
	authRepo := model.AuthRepo{}
	if err := r.getItem(ctx, entity.Auth, &authRepo, email); err != nil {
		return model.Auth{}, err
	}

	return authRepo.ToAPI(), nil
}
