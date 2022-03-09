package client

import (
	"context"
	"github.com/reb-felipe/grpc-from-rest/domain/entity"
)

type UsersClient interface {
	CreateUser(ctx context.Context, name string, coordinates []float64) (*entity.User, error)
	UpdateUser(ctx context.Context, userID string, name string, coordinates []float64) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
}
