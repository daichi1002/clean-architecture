package server

import (
	"clean-architecture/infra"
	"clean-architecture/pb"
	"clean-architecture/server/interceptor"
	"clean-architecture/server/service"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

func NewGRPCServer(gormHandler *infra.GormHandler) *grpc.Server {
	options := make([]grpc.ServerOption, 0)
	options = append(options, grpc.UnaryInterceptor(
		middleware.ChainUnaryServer(
			interceptor.ValidationInterceptor(),
		)),
	)

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, service.NewUserService(gormHandler))

	return server
}
