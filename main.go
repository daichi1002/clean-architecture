package main

import (
	"clean-architecture/infra"
)

func main() {
	// router初期化処理
	infra.InitRouter()
	// API接続
	infra.Router.Run()
}
