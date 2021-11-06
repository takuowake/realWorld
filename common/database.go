package common

import "github.com/jinzhu/gorm"

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
