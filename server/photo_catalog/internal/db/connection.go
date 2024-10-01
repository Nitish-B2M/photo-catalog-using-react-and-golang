package db

import (
	"github.com/photo_catalog/internal/modals"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnection() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3307)/photo_catalog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&modals.CatalogItem{}); err != nil {
		return nil, err
	}

	return db, nil
}
