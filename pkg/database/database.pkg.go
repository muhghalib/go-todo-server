package database

import (
	"fmt"
	"log"
	"main/internal/entities"
	"main/pkg/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(config config.Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Task{})

	connection, err := db.DB()

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(20)
	connection.SetMaxOpenConns(100)
	connection.SetConnMaxLifetime(time.Hour)

	return db
}
