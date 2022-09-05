package infra

import (
	gin "github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"clean-architecture/constant"
	"clean-architecture/interface/controller"
)

var Router *gin.Engine

func InitRouter() {
	// 環境変数読み込み
	loadEnv()

	router := gin.Default()
	gormHandler := NewGormHandler()

	userController := controller.NewUserController(gormHandler)
	// エンドポイント
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.POST("/user", func(c *gin.Context) { userController.Create(c) })

	Router = router
}

func loadEnv() {
	viper.AutomaticEnv()
	viper.SetDefault(constant.DBHostEnv, "127.0.0.1")
	viper.SetDefault(constant.DBPortEnv, "3306")
	viper.SetDefault(constant.DBUserEnv, "root")
	viper.SetDefault(constant.DBPasswordEnv, "daichi1002")
	viper.SetDefault(constant.DBNameEnv, "dev")
}
