# PRACTICE GO Hexagonal Architecture

## Request Implement

### structure

- go-crud-jwt_auth_cookie-hexagonal
  - main.go
  - adapters
     - gorm_adapter_test.go
     - gorm.book.go
     - gorm.go
     - gorm.user.go
     - http_adapter_test.go
     - http.book.go
     - http.user.go
  - core
     - book.repository.go
     - book.service.go
     - books.go
     - user_service.test.go
     - user.go
     - user.repository.go
     - user.service.go

### user

- POST    /register         create a new user
- POST    /login            login method
- PUT     /user/edit        update user
- DELETE  /user/delete      delete user

### book

- GET     /books            get all book
- GET     /books/:name      get all the same book
- POST    /books            create new book
- POST    /books/author     create new author
- POST    /books/publisher  create new publisher
- PUT     /books/:id        update book by id
- DELETE  /books/:id        delete book by id

## About

- DB: PostgreSQL
- ORM: gorm
- HTTP: fiber

### HOW TO START

docker-compose up -d
