package repositories

import (
	"BookShopGoProject/internal/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type IBookRepo interface {
	GetBooks() ([]models.Book, error)
	GetBookById(id uint) (models.Book, error)
	GetBookByName(name string) (*models.Book, error)
	CreateBook(book *models.Book, quantity uint) error
}

type BookRepo struct {
	DB *gorm.DB
}

func (b *BookRepo) GetBookById(id uint) (models.Book, error) {
	var book models.Book
	err := b.DB.Preload("Authors").Preload("Stocks").First(&book, id).Error
	return book, err
}

func (b *BookRepo) GetBooks() ([]models.Book, error) {
	if b.DB == nil {
		log.Fatal("Received nil database instance in GetBooks")
	}
	var books []models.Book
	err := b.DB.Preload("Authors").Preload("Stocks").Find(&books).Error
	return books, err
}

func (b *BookRepo) GetBookByName(name string) (*models.Book, error) {
	var book models.Book
	err := b.DB.Debug().Where("book_name = ?", name).Preload("Authors").Preload("Stocks").First(&book).Error

	return &book, err
}

func (b *BookRepo) CreateBook(book *models.Book, quantity uint) error {
	err := b.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Select("BookName", "Description", "Price", "Genre", "Status").Create(&book).Error
		if err != nil {
			return err
		}
		fmt.Println("Book Created: ", book)
		var bookAuthorList []models.BookAuthor
		for i := range book.Authors {
			bookAuthor := models.BookAuthor{
				BookId:   book.BookId,
				AuthorId: book.Authors[i].AuthorId,
			}
			bookAuthorList = append(bookAuthorList, bookAuthor)
		}
		err = tx.Debug().Create(&bookAuthorList).Error
		if err != nil {
			fmt.Printf("Can't insert book authors: %v\n", err)
			return err
		}
		stock := models.Stock{
			BookId:   book.BookId,
			Quantity: quantity,
			UpdateAt: time.Now(),
		}
		err = tx.Debug().Create(&stock).Error
		if err != nil {
			fmt.Printf("Can't insert stocks: %v\n", err)
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
