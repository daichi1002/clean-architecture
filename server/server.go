package server

import (
	"clean-architecture/infra"
	"clean-architecture/pb"
	"clean-architecture/server/service"

	"google.golang.org/grpc"
)

func NewGRPCServer(gormHandler *infra.GormHandler) *grpc.Server {
	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, service.NewUserService(gormHandler))

	return server
}
