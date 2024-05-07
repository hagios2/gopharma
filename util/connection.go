package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	dsn := "user:pass@tcp(127.0.0.1:3306)/gopharma?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		err = fmt.Errorf("could not connect to dB")
		return
	}

	DB = db
}
