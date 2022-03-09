package server

import (
	"context"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/grpc/presenter"
	gen "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UsersController interface {
	Create(ctx context.Context, input *presenter.UserInput) (*presenter.User, error)
	Update(ctx context.Context, input *presenter.UpdateUserInput) (*presenter.User, error)
	Delete(ctx context.Context, input *presenter.IDInput) (*emptypb.Empty, error)
	Get(ctx context.Context, input *presenter.IDInput) (*presenter.User, error)
	List(ctx context.Context, empty *emptypb.Empty) (*presenter.UsersList, error)
}

type usersServer struct {
	*gen.UnimplementedUsersServer
	UsersController
}

func newServer(controller UsersController) *usersServer {
	return &usersServer{
		UnimplementedUsersServer: new(gen.UnimplementedUsersServer),
		UsersController:          controller,
	}
}

func (g *usersServer) Create(ctx context.Context, input *presenter.UserInput) (*presenter.User, error) {
	return g.UsersController.Create(ctx, input)
}

func (g *usersServer) Update(ctx context.Context, input *presenter.UpdateUserInput) (*presenter.User, error) {
	return g.UsersController.Update(ctx, input)
}

func (g *usersServer) Delete(ctx context.Context, input *presenter.IDInput) (*emptypb.Empty, error) {
	return g.UsersController.Delete(ctx, input)
}

func (g *usersServer) Get(ctx context.Context, input *presenter.IDInput) (*presenter.User, error) {
	return g.UsersController.Get(ctx, input)
}

func (g *usersServer) List(ctx context.Context, empty *emptypb.Empty) (*presenter.UsersList, error) {
	return g.UsersController.List(ctx, empty)
}
