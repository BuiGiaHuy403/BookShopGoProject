package models

import "time"

type Order struct {
	OrderId    uint      `gorm:"primaryKey;autoIncrement" json:"order_id"`
	TotalPrice float64   `gorm:"column:total_price" json:"total_price"`
	Status     string    `gorm:"column:status" json:"status"`
	OrderDate  time.Time `gorm:"column:order_date" json:"order_date"`

	Books []Book `gorm:"many2many:order_books;joinForeignKey:order_id;joinReferences:book_id" json:"books"`
}
