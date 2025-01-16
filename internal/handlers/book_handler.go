package handlers

import (
	"BookShopGoProject/internal/models"
	"BookShopGoProject/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	IBookService services.IBookService
}

func (b *BookHandler) GetBooks(c *gin.Context) {
	var name = c.Query("name")
	if name != "" {
		book, err := b.IBookService.GetBookByName(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
		c.JSON(http.StatusOK, book)
		return
	} else {
		books, err := b.IBookService.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
		c.JSON(http.StatusOK, books)
		return
	}

}

func (b *BookHandler) GetBookById(c *gin.Context) {
	var data = c.Param("id")
	id, err := strconv.Atoi(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	book, err := b.IBookService.GetBookById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	c.JSON(http.StatusOK, book)
}

func (b *BookHandler) CreateBook(c *gin.Context) {
	var newBook models.BookCreation
	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = newBook.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var book = models.Book{
		BookName:    newBook.BookName,
		Description: newBook.Description,
		Price:       newBook.Price,
		Genre:       newBook.Genre,
		Authors:     newBook.Authors,
		Status:      true,
	}

	var quantity = newBook.Quantity

	err = b.IBookService.CreateBook(&book, quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book":             book,
		"response_message": "Book created successfully",
		"status_code":      http.StatusOK,
	})
}
