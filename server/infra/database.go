package infra

import (
	"fmt"
	"log"

	"clean-architecture/constant"
	"clean-architecture/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormHandler struct {
	DB *gorm.DB
}

func NewGormHandler() *GormHandler {
	conn, err := connectDB()
	if err != nil {
		panic(err.Error)
	}
	GormHandler := new(GormHandler)
	GormHandler.DB = conn
	return GormHandler
}

func connectDB() (*gorm.DB, error) {

	dbHost := viper.GetString(constant.DBHostEnv)
	dbPort := viper.GetString(constant.DBPortEnv)
	dBName := viper.GetString(constant.DBNameEnv)
	dbUser := viper.GetString(constant.DBUserEnv)
	dbPassword := viper.GetString(constant.DBPasswordEnv)

	dbAuth := dbUser
	if len(dbPassword) > 0 {
		dbAuth += ":" + dbPassword
	}
	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=10s", dbAuth, dbHost, dbPort, dBName)
	var db *gorm.DB
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrateErr := db.AutoMigrate(&domain.User{})

	if migrateErr != nil {
		log.Fatal(err)
	}

	return db, nil
}

func (handler *GormHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.DB.Find(out, where...)
}

func (handler *GormHandler) Store(obj interface{}) *gorm.DB {
	return handler.DB.Create(obj)
}
