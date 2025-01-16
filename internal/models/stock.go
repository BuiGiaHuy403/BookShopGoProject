package models

import "time"

type Stock struct {
	StockId  uint      `gorm:"primaryKey;autoIncrement" json:"stock_id"`
	Quantity uint      `gorm:"column:quantity" json:"quantity"`
	UpdateAt time.Time `gorm:"column:update_at" json:"update_at"`

	BookId uint `gorm:"foreignKey;column:book_id" json:"book_id"`
}
