package usecase

import (
	"clean-architecture/domain"
	"clean-architecture/util"
	"time"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Users() (user *domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) CreateUser(u *domain.User) (user *domain.User, err error) {
	userID, err := util.NewULID()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user = &domain.User{
		UserId:    userID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Birthday:  now,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = interactor.UserRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor *UserInteractor) ShowUser(id string) (user *domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)

	if user.UserId == "" {
		return nil, nil
	}

	return
}

func (interactor *UserInteractor) Test() {}
