package main

import (
	"github.com/reb-felipe/grpc-from-rest/cmd/client"
	wrapper "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/rest/client"
)

func main() {
	terminal, err := client.NewTerminal()
	if err != nil {
		panic(err)
	}

	usersClient := wrapper.NewClient("http://localhost:8081/users")
	app := client.NewApp(terminal, usersClient)
	app.Run()
}
