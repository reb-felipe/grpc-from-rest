proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/*.proto

env:
	go mod vendor

rest-server:
	go run cmd/server/rest/main.go

rest-client:
	go run cmd/client/rest/main.go

grpc-server:
	go run cmd/server/grpc/main.go

grpc-client:
	go run cmd/client/grpc/main.go