package database

import (
	"fmt"
	"log"
	"rizkiwhy-todo-app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MySQLConnection() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("database connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	return
}
