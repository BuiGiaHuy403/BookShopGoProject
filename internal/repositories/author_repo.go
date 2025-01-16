package repositories

import (
	"gorm.io/gorm"
)

type IAuthorRepo interface {
	GetAuthorById()
}

type AuthorRepo struct {
	db *gorm.DB
}

//
//func (repo *AuthorRepo) GetAuthorById() []models.Author {
//	repo.db = config.DB
//
//}
