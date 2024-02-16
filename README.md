# PRACTICE GO Hexagonal Architecture

## Request Implement

### user

- POST    /register         create new user
- POST    /login            login method
- PUT     /user/edit        update user
- DELETE  /user/delete      delete user

### book

- GET     /books            get all book
- GET     /books/:name      get all same book
- POST    /books            create new book
- POST    /books/author     create new author
- POST    /books/publisher  create new publisher
- PUT     /books/:id        update book by id
- DELETE  /books/:id        delete book by id

## About

- DB: postgresql
- ORM: gorm
- HTTP: fiber

### HOW TO START

docker-compose up -d
