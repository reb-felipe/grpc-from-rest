package main

import (
	"github.com/reb-felipe/grpc-from-rest/domain/service"
	"github.com/reb-felipe/grpc-from-rest/internal/controller/http/grpc"
	"github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc/server"
	"github.com/reb-felipe/grpc-from-rest/internal/infrastructure/storage/mem"
	"log"
	"net"
)

func main() {
	repo := mem.NewUsers()
	svc := service.NewUsers(repo)
	ctrl := grpc.NewUsers(svc)
	srv := server.NewServer(ctrl)

	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	log.Fatal(srv.Serve(l))
}
