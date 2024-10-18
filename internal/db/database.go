package db

import (
	"fmt"
	"log"

	"github.com/photo_catalog/internal/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3307)/photo_catalog?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("DB:", DB)
	// Auto Migrate
	DB.AutoMigrate(&entities.CatalogItem{})
	DB.AutoMigrate(&entities.UserRegister{})
	return DB
}
