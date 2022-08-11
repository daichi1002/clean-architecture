package database

import "gorm.io/gorm"

type GormHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
}
