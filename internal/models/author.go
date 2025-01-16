package models

type Author struct {
	AuthorId   uint   `gorm:"primaryKey;AutoIncrement;column:author_id;" json:"author_id"`
	AuthorName string `gorm:"column:author_name" json:"author_name"`
	Status     bool   `gorm:"column:status" json:"status"`

	Books []Book `gorm:"many2many:book_authors;joinForeignKey:author_id;joinReferences:book_id" json:"books"`
}
