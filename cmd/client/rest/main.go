package main

import (
	"github.com/reb-felipe/grpc-from-rest/cmd/client"
	rest "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/rest/client"
)

func main() {
	t, err := client.NewTerminal()
	if err != nil {
		panic(err)
	}
	cl := client.NewApp(t, rest.NewClient("http://localhost:8081/users"))
	cl.Run()
}
