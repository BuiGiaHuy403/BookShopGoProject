package models

type OrderBook struct {
	OrderId  uint `gorm:"primaryKey" json:"order_id"`
	BookId   uint `gorm:"primaryKey" json:"book_id"`
	Quantity uint `gorm:"column:quantity" json:"quantity"`
}
