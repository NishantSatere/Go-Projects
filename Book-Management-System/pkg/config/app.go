package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:nishant@tcp(localhost:3306)/goProject?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
