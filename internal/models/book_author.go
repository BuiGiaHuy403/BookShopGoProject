package models

type BookAuthor struct {
	AuthorId uint `gorm:"primaryKey" json:"author_id"`
	BookId   uint `gorm:"primaryKey" json:"book_id"`
}
