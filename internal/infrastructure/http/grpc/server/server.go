package server

import (
	gen "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc"
	"google.golang.org/grpc"
)

func NewServer(usersController UsersController) *grpc.Server {
	usersServer := newServer(usersController)

	mainServer := grpc.NewServer()
	gen.RegisterUsersServer(mainServer, usersServer)
	return mainServer
}
