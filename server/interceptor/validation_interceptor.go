package interceptor

import (
	"clean-architecture/server/validator"
	"context"

	"google.golang.org/grpc"
)

func ValidationInterceptor() grpc.UnaryServerInterceptor {
	userValidator := validator.NewUserValidator()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ok, err := userValidator.Validate(ctx, req)
		if ok && err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
