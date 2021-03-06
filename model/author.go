package model

import (
	"gorm.io/gorm"
)
//Model for Author
type Author struct {
	gorm.Model //provides CreatedAt and UpdatedAt
	Name	string `json:"name"`
	// Books []Book //author has many books associated
}

type AuthorService interface {
	Create(author *Author) (*Author, error)
	FindAll() ([]Author, error)
	Delete(author *Author) error
	FindById(id int64) ([]Author, error)
	GetTotalNumberOfAuthors() (int64, error)
	GetAuthorsNameList() ([]string, error) 
}

type AuthorRepository interface {
	Save(author *Author) (*Author, error)
	FindAll() ([]Author, error)
	Delete (author *Author) error
	FindById(id int64) ([]Author, error)
	TotalNumberOfAuthors() (int64, error)
	AllAuthorsNameList() ([]string, error)
	Migrate() error
}