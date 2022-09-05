package usecase

import "clean-architecture/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Users() (user *domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) CreateUser(u *domain.User) (err error) {
	err = interactor.UserRepository.Create(u)
	return
}
