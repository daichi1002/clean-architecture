package controller

import (
	"clean-architecture/interface/database"
	"clean-architecture/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

// controllerの初期化
func NewUserController(gormHandler database.GormHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				GormHandler: gormHandler,
			},
		},
	}
}

// ユーザー一覧取得
func (controller *UserController) Index(c Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}
