package main

import (
	"github.com/reb-felipe/grpc-from-rest/domain/service"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/rest"
	"github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/rest/server"
	"github.com/reb-felipe/grpc-from-rest/internal/infrastructure/storage/mem"
	"log"
)

func main() {
	repo := mem.NewUsers()
	svc := service.NewUsers(repo)
	ctrl := rest.NewRest(svc)
	srv := server.NewServer(&server.Config{
		Controller: ctrl,
		Addr:       ":8081",
	})
	log.Fatal(srv.ListenAndServe())
}
