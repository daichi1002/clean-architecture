package server

import (
	"clean-architecture/infra"
	"clean-architecture/pb"
	"clean-architecture/server/service"

	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	gormHandler := infra.NewGormHandler()
	pb.RegisterUserServiceServer(server, service.NewUserService(gormHandler))

	return server
}
