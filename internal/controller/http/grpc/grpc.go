package grpc

import (
	"context"
	"github.com/reb-felipe/grpc-from-rest/domain/service"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/grpc/presenter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewUsers(users *service.Users) *User {
	return &User{users: users}
}

type User struct {
	users *service.Users
}

func (u *User) Create(ctx context.Context, input *presenter.UserInput) (*presenter.User, error) {
	user, err := u.users.CreateUser(ctx, input.ToEntity())
	if err != nil {
		return nil, err
	}

	return presenter.UserFromEntity(user), nil
}

func (u *User) Update(ctx context.Context, input *presenter.UpdateUserInput) (*presenter.User, error) {
	user, err := u.users.UpdateUser(ctx, input.ToEntity())
	if err != nil {
		return nil, err
	}

	return presenter.UserFromEntity(user), nil
}

func (u *User) Delete(ctx context.Context, input *presenter.IDInput) (*emptypb.Empty, error) {
	return new(emptypb.Empty), u.users.DeleteUser(ctx, input.Id)
}

func (u *User) Get(ctx context.Context, input *presenter.IDInput) (*presenter.User, error) {
	user, err := u.users.GetUser(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	return presenter.UserFromEntity(user), nil
}

func (u *User) List(ctx context.Context, empty *emptypb.Empty) (*presenter.UsersList, error) {
	users, err := u.users.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*presenter.User, len(users))
	for i, v := range users {
		result[i] = presenter.UserFromEntity(v)
	}

	return &presenter.UsersList{Users: result}, nil
}
