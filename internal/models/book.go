package models

import (
	"errors"
)

type Book struct {
	BookId      uint     `gorm:"primaryKey;autoIncrement;column:book_id" json:"book_id"`
	BookName    string   `gorm:"column:book_name" json:"book_name"`
	Description string   `gorm:"column:description" json:"description"`
	Price       float64  `gorm:"column:price" json:"price"`
	Genre       string   `gorm:"column:genre" json:"genre"`
	Status      bool     `gorm:"column:status" json:"status"`
	Authors     []Author `gorm:"many2many:book_authors;joinForeignKey:book_id;joinReferences:author_id" json:"authors"`
	Orders      []Order  `gorm:"many2many:order_books;joinForeignKey:book_id;joinReferences:order_id" json:"orders"`
	Stocks      []Stock  `gorm:"foreignKey:book_id" json:"stocks"`
}

type BookCreation struct {
	BookName    string   ` json:"book_name"`
	Description string   ` json:"description"`
	Price       float64  ` json:"price"`
	Genre       string   ` json:"genre"`
	Authors     []Author ` json:"authors"`
	Quantity    uint     ` json:"quantity"`
}

func (b *BookCreation) Validate() error {
	if b.BookName == "" {
		return errors.New("Book name is required")
	}
	if b.Description == "" {
		return errors.New("Description is required")
	}
	if b.Price <= 0 {
		return errors.New("Price must be greater than zero")
	}
	if b.Genre == "" {
		return errors.New("Genre is required")
	}
	if len(b.Authors) == 0 {
		return errors.New("At least one author is required")
	}
	if b.Quantity < 1 {
		return errors.New("Quantity must be greater than zero")
	}
	return nil
}

type BookRepoInput struct {
	Book     *Book
	Quantity uint
}
