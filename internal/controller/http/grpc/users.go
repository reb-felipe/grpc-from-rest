package grpc

import (
	"context"
	"github.com/reb-felipe/grpc-from-rest/domain/service"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/grpc/presenter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewUsers(users *service.Users) *UsersController {
	return &UsersController{usersService: users}
}

type UsersController struct {
	usersService *service.Users
}

func (u *UsersController) Create(ctx context.Context, input *presenter.UserInput) (*presenter.User, error) {
	user, err := u.usersService.CreateUser(ctx, input.ToEntity())
	if err != nil {
		return nil, err
	}

	return presenter.UserFromEntity(user), nil
}

func (u *UsersController) Update(ctx context.Context, input *presenter.UpdateUserInput) (*presenter.User, error) {
	user, err := u.usersService.UpdateUser(ctx, input.ToEntity())
	if err != nil {
		return nil, err
	}

	return presenter.UserFromEntity(user), nil
}

func (u *UsersController) Delete(ctx context.Context, input *presenter.IDInput) (*emptypb.Empty, error) {
	return new(emptypb.Empty), u.usersService.DeleteUser(ctx, input.Id)
}

func (u *UsersController) Get(ctx context.Context, input *presenter.IDInput) (*presenter.User, error) {
	user, err := u.usersService.GetUser(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	return presenter.UserFromEntity(user), nil
}

func (u *UsersController) List(ctx context.Context, empty *emptypb.Empty) (*presenter.UsersList, error) {
	users, err := u.usersService.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*presenter.User, len(users))
	for i, v := range users {
		result[i] = presenter.UserFromEntity(v)
	}

	return &presenter.UsersList{Users: result}, nil
}
