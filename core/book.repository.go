package core

// PORT FOR SECONDARY ADAPTERS
type BookRepository interface {
	CreateBook(book Book) error
	CreatePublisher(publisher Publisher) error
	CreateAuthor(author Author) error
	ReadBooks() ([]Book, error)
	ReadNameBook(bookName string) ([]Book, error)
	UpdateBook(bookId int, book Book) error
	DeleteBook(bookId int) error
}
