package main

import (
	"github.com/reb-felipe/grpc-from-rest/cmd/client"
	gen "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc"
	wrapper "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	terminal, err := client.NewTerminal()
	if err != nil {
		panic(err)
	}

	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	grpcUserClient := gen.NewUsersClient(connection)

	usersClient := wrapper.NewClient(grpcUserClient)
	app := client.NewApp(terminal, usersClient)
	app.Run()
}
