package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:root@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	db = d
	log.Println("Connected to MySQL successfuly....")
}

func GetDB() *gorm.DB {
	return db
}
