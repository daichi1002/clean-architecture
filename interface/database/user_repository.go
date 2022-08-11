package database

import (
	"clean-architecture/domain"
)

type UserRepository struct {
	GormHandler
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	if err = repo.Find(&users).Error; err != nil {
		return
	}
	return
}
