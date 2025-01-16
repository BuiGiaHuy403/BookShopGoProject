package repositories

import (
	"BookShopGoProject/internal/models"
	"gorm.io/gorm"
)

func CreateStock(stock models.Stock, db *gorm.DB) error {
	if err := db.Create(stock).Error; err != nil {
		return err
	}
	return nil
}
