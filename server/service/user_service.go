package service

import (
	"clean-architecture/domain"
	"clean-architecture/interface/database"
	"clean-architecture/pb"
	"clean-architecture/usecase"
	"context"
	"fmt"
)

type UserService struct {
	Interactor usecase.UserInteractor
}

// NewUserService is create new instance of pb.UserServiceServer.
func NewUserService(gormHandler database.GormHandler) *UserService {
	return &UserService{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				GormHandler: gormHandler,
			},
		},
	}
}

// ユーザー一覧取得
func (service *UserService) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := service.Interactor.Users()
	if err != nil {
		fmt.Println(err)
	}
	return &pb.ListUsersResponse{
		Users: users.ToProto(),
	}, nil
}

// ユーザー作成
func (service *UserService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	param := domain.User{
		Name:  request.Name,
		Email: request.Email,
	}

	user, err := service.Interactor.CreateUser(&param)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: user.ToProto()}, nil
}

// ユーザー取得
