package main

//import (
//	"github.com/reb-felipe/grpc-from-rest/cmd/client"
//	grpc2 "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc"
//	grpc3 "github.com/reb-felipe/grpc-from-rest/internal/infrastructure/http/grpc/client"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//)
//
//func main() {
//	t, err := client.NewTerminal()
//	if err != nil {
//		panic(err)
//	}
//
//	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		panic(err)
//	}
//	cl := client.NewApp(t, grpc3.NewClient(grpc2.NewUsersClient(conn)))
//	cl.Run()
//}
