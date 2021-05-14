package config

import (
	"os"

	_ "github.com/jinzhu/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
)

var db *gorm.DB

func init() {
	gotenv.Load("../.env")
}

func Connect() {
	d, err := gorm.Open("mysql", os.Getenv("DB_STRING"))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
