package adapters

import (
	core "github.com/hamwiwatsapon/go-crud-authen/core"
	"gorm.io/gorm"
)

func NewGormBook(db *gorm.DB) core.BookRepository {
	return &GormBookStore{db: db}
}

// CreateBook implements core.BookRepository.
func (r *GormBookStore) CreateBook(book core.Book) error {
	if result := r.db.Create(&book); result.Error != nil {
		return result.Error
	}
	return nil
}

// ReadBooks implements core.BookRepository.
func (r *GormBookStore) ReadBooks() ([]core.Book, error) {
	books := new([]core.Book)
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return *books, nil
}

// ReadBook implements core.BookRepository.
func (r *GormBookStore) ReadNameBook(bookName string) ([]core.Book, error) {
	books := new([]core.Book)
	result := r.db.Where(&core.Book{Name: bookName}).Find(books)

	if result.Error != nil {
		return nil, result.Error
	}

	return *books, nil
}

// Update implements core.BookRepository.
func (r *GormBookStore) UpdateBook(bookId int, book core.Book) error {
	result := r.db.Model(core.Book{}).Where("id = ?", bookId).Updates(core.Book{
		Name:        book.Name,
		Author:      book.Author,
		Description: book.Description,
		PublisherID: book.PublisherID,
		GenreID:     book.GenreID})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Delete implements core.BookRepository.
func (r *GormBookStore) DeleteBook(bookId int) error {
	selectedBook := new(core.Book)
	result := r.db.Where("id = ?", bookId).Delete(selectedBook)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CreateAuthor implements core.BookRepository.
func (r *GormBookStore) CreateAuthor(author core.Author) error {
	if result := r.db.Create(&author); result.Error != nil {
		return result.Error
	}
	return nil
}

// CreatePublisher implements core.BookRepository.
func (r *GormBookStore) CreatePublisher(publisher core.Publisher) error {
	if result := r.db.Create(&publisher); result.Error != nil {
		return result.Error
	}
	return nil
}
