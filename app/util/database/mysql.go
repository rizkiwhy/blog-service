package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"rizkiwhy-blog-service/util/config"
)

func MySQLConnection() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env("DB_USER"),
		config.Env("DB_PASSWORD"),
		config.Env("DB_HOST"),
		config.Env("DB_PORT"),
		config.Env("DB_NAME"))

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		return nil, err
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	return
}
