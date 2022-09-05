package database

import (
	"clean-architecture/domain"
)

type UserRepository struct {
	GormHandler
}

func (repo *UserRepository) FindAll() (users *domain.Users, err error) {
	err = repo.Find(&users).Error
	return
}

func (repo *UserRepository) Create(user *domain.User) (err error) {
	err = repo.Store(&user).Error
	return
}
