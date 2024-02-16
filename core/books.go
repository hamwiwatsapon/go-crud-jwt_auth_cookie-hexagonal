package core

import "gorm.io/gorm"

/*
	Book Models
*/
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`

	// Publisher FK
	PublisherID uint
	Publisher   Publisher

	// Genre FK
	GenreID uint
	Genre   Genre

	Authors []Author `gorm:"many2many:author_books;"`
}

type Genre struct {
	gorm.Model
	Details string
	Name    string
}

type Publisher struct {
	gorm.Model
	Details string
	Name    string
}

type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"many2many:author_books;"`
}

type AuthorBook struct {
	AuthorID uint
	Author   Author
	BookID   uint
	Book     Book
}
