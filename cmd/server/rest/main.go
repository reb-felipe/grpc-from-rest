package main

import (
	"github.com/reb-felipe/grpc-from-rest/domain/service"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/rest"
	"github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/rest/server"
	"github.com/reb-felipe/grpc-from-rest/internal/infrastructure/storage/mem"
	"log"
)

func main() {
	repository := mem.NewUsers()
	service := service.NewUsers(repository)
	controller := rest.NewUsers(service)
	server := server.NewServer(&server.Config{
		Controller: controller,
		Addr:       ":8081",
	})
	log.Fatal(server.ListenAndServe())
}
