package server

import (
	"context"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/grpc/presenter"
	grpc2 "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type Server interface {
	Serve(net.Listener) error
	Stop()
	GracefulStop()
}

type Controller interface {
	Create(ctx context.Context, input *presenter.UserInput) (*presenter.User, error)
	Update(ctx context.Context, input *presenter.UpdateUserInput) (*presenter.User, error)
	Delete(ctx context.Context, input *presenter.IDInput) (*emptypb.Empty, error)
	Get(ctx context.Context, input *presenter.IDInput) (*presenter.User, error)
	List(ctx context.Context, empty *emptypb.Empty) (*presenter.UsersList, error)
}

func newServer(controller Controller) *grpcServer {
	return &grpcServer{
		UnimplementedUsersServer: new(grpc2.UnimplementedUsersServer),
		Controller:               controller,
	}
}

type grpcServer struct {
	*grpc2.UnimplementedUsersServer
	Controller
}

func (g *grpcServer) Create(ctx context.Context, input *presenter.UserInput) (*presenter.User, error) {
	return g.Controller.Create(ctx, input)
}

func (g *grpcServer) Update(ctx context.Context, input *presenter.UpdateUserInput) (*presenter.User, error) {
	return g.Controller.Update(ctx, input)
}

func (g *grpcServer) Delete(ctx context.Context, input *presenter.IDInput) (*emptypb.Empty, error) {
	return g.Controller.Delete(ctx, input)
}

func (g *grpcServer) Get(ctx context.Context, input *presenter.IDInput) (*presenter.User, error) {
	return g.Controller.Get(ctx, input)
}

func (g *grpcServer) List(ctx context.Context, empty *emptypb.Empty) (*presenter.UsersList, error) {
	return g.Controller.List(ctx, empty)
}

func NewServer(controller Controller) Server {
	service := newServer(controller)
	server := grpc.NewServer()

	grpc2.RegisterUsersServer(server, service)
	return server
}
