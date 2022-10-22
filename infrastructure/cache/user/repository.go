package user

import (
	"context"

	"github.com/buley/n3ut/domain"
	"github.com/buley/n3ut/domain/user"
	"github.com/buley/n3ut/library/kvs"
)

type repository struct{}

func NewRepository() user.Repository {
	return &repository{}
}

func (repo *repository) Add(ctx context.Context, u *domain.User) error {
	if _, ok := kvs.Get(u.Address.Hex()); ok {
		return domain.ErrUserAlreadyExists
	}

	kvs.Set(u.Address.Hex(), u)

	return nil
}

func (repo *repository) Get(ctx context.Context, address domain.Address) (*domain.User, error) {
	data, ok := kvs.Get(address.Hex())
	if !ok {
		return nil, domain.ErrUserNotFound
	}

	u, ok := data.(*domain.User)
	if !ok {
		return nil, domain.ErrUserBroken
	}

	return u, nil
}

func (repo *repository) Update(ctx context.Context, u *domain.User) error {
	if _, ok := kvs.Get(u.Address.Hex()); !ok {
		return domain.ErrUserNotFound
	}

	kvs.Set(u.Address.Hex(), u)

	return nil
}

func (repo *repository) Delete(ctx context.Context, u *domain.User) error {
	if _, ok := kvs.Delete(u.Address.Hex()); !ok {
		return domain.ErrUserNotFound
	}

	return nil
}
