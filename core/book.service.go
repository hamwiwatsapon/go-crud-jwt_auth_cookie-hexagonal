package core

import (
	"errors"
	"net/url"
)

type BookService interface {
	NewBook(book Book) error
	ReadBooks() ([]Book, error)
	ReadNameBook(bookName string) ([]Book, error)
	UpdateBook(bookId int, book Book) error
	DeleteBook(bookId int) error
}

type BookServiceImpl struct {
	repo BookRepository
}

func (s *BookServiceImpl) NewBook(book Book) error {
	if book.Name == "" {
		return errors.New("name cant be empty string")
	}

	if err := s.repo.CreateBook(book); err != nil {
		return err
	}

	return nil
}

func (s *BookServiceImpl) ReadBooks() ([]Book, error) {
	books, err := s.repo.ReadBooks()
	if err != nil {
		return books, err
	}
	return books, err
}

func (s *BookServiceImpl) ReadNameBook(bookName string) ([]Book, error) {
	decodedBookName, err := url.QueryUnescape(bookName)

	if err != nil {
		return nil, err
	}

	books, err := s.repo.ReadNameBook(decodedBookName)
	if err != nil {
		return nil, err
	}
	return books, err
}

func (s *BookServiceImpl) UpdateBook(bookId int, book Book) error {
	err := s.repo.UpdateBook(bookId, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookServiceImpl) DeleteBook(bookId int) error {
	err := s.repo.DeleteBook(bookId)
	if err != nil {
		return err
	}
	return nil
}

func NewBookService(repo BookRepository) BookService {
	return &BookServiceImpl{repo: repo}
}
