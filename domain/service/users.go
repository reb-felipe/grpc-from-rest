package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
	"time"
)

type Users struct {
	repository UsersRepository
}

func (u *Users) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()
	return u.repository.CreateUser(ctx, user)
}

func (u *Users) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	now := time.Now()
	user.UpdatedAt = &now
	return u.repository.UpdateUser(ctx, user)
}

func (u *Users) ListUsers(ctx context.Context) ([]*entity.User, error) {
	return u.repository.ListUsers(ctx)
}

func (u *Users) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	return u.repository.GetUser(ctx, userID)
}

func (u *Users) DeleteUser(ctx context.Context, userID string) error {
	return u.repository.DeleteUser(ctx, userID)
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	DeleteUser(ctx context.Context, userID string) error
}
