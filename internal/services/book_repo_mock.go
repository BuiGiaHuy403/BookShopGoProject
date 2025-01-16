package services

import (
	"BookShopGoProject/internal/models"
	"github.com/stretchr/testify/mock"
)

type BookRepoMock struct {
	mock.Mock
}

func (b *BookRepoMock) GetBooks() ([]models.Book, error) {
	args := b.Called()
	r0 := args.Get(0)
	var v0 []models.Book
	if r0 != nil {
		v0 = r0.([]models.Book)
	}
	r1 := args.Get(1)
	var v1 error
	if r1 != nil {
		v1 = r1.(error)
	}
	return v0, v1
}

func (b *BookRepoMock) GetBookById(id uint) (models.Book, error) {
	args := b.Called(id)
	v0 := args.Get(0).(models.Book)
	r1 := args.Get(1)
	var v1 error
	if r1 != nil {
		v1 = r1.(error)
	}
	return v0, v1
}

func (b *BookRepoMock) GetBookByName(name string) (*models.Book, error) {
	args := b.Called(name)
	r0 := args.Get(0)
	var v0 *models.Book
	if r0 != nil {
		v0 = r0.(*models.Book)
	}
	r1 := args.Get(1)
	var v1 error
	if r1 != nil {
		v1 = r1.(error)
	}
	return v0, v1
}

func (b *BookRepoMock) CreateBook(book *models.Book, quantity uint) error {
	args := b.Called(book, quantity)
	r0 := args.Get(0)
	var v0 error
	if r0 != nil {
		v0 = r0.(error)
	}
	return v0
}
