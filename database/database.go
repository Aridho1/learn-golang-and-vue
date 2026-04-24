package database

import (
	"fmt"
	"log"

	"zunn/backend-api/config"
	"zunn/backend-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER", "root"),
		config.GetEnv("DB_PASS", ""),
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_PORT", "3306"),
		config.GetEnv("DB_NAME", ""),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		// &models.Product{},
	)
	if err != nil {
		log.Fatal("Failed migrate:", err)
	}

	fmt.Println("Database connected & migrated")
}
