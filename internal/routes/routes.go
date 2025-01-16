package routes

import (
	"BookShopGoProject/internal/handlers"
	"BookShopGoProject/internal/repositories"
	"BookShopGoProject/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, DB *gorm.DB) {
	v1 := router.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			BookHandler := handlers.BookHandler{
				IBookService: &services.BookService{
					IBookRepo: &repositories.BookRepo{
						DB: DB,
					},
				},
			}

			books.GET("", BookHandler.GetBooks)
			books.GET(":id", BookHandler.GetBookById)
			books.POST("", BookHandler.CreateBook)
		}
	}
}
