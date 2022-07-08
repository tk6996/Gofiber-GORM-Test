package database

import (
	"fmt"
	"go-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		"root",
		"ADMIN1234",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DBConn.AutoMigrate(&models.User{})
}
