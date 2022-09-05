package main

import (
	"clean-architecture/infra"
	"clean-architecture/interface/controller"
)

func main() {
	// router初期化処理
	infra.InitRouter()
	// API接続
	err := infra.Router.Run()
	if err != nil {
		controller.NewError(err)
	}
}
