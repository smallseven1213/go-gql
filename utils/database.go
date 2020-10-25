package utils

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sys?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
	}

	// db.DB().SetMaxIdleConns(10)
	// db.LogMode(true)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
