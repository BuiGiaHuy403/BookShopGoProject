package services

import (
	"BookShopGoProject/internal/models"
	"BookShopGoProject/internal/repositories"
	"errors"
)

type IBookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookById(bookId uint) (models.Book, error)
	CreateBook(book *models.Book, quantity uint) error
	GetBookByName(bookName string) (*models.Book, error)
}

type BookService struct {
	IBookRepo repositories.IBookRepo
}

func (BookService *BookService) GetBookByName(bookName string) (*models.Book, error) {
	return BookService.IBookRepo.GetBookByName(bookName)
}

func (BookService *BookService) GetAllBooks() ([]models.Book, error) {
	return BookService.IBookRepo.GetBooks()
}

func (BookService *BookService) GetBookById(id uint) (models.Book, error) {
	return BookService.IBookRepo.GetBookById(id)
}

func (BookService *BookService) CreateBook(book *models.Book, quantity uint) error {

	existingBook, err := BookService.IBookRepo.GetBookByName(book.BookName)
	if existingBook.BookId != 0 {
		return errors.New("book already exists")
	}

	err = BookService.IBookRepo.CreateBook(book, quantity)
	if err != nil {
		return errors.New("failed to create book")
	}

	return nil

}
