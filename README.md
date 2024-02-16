# PRACTICE GO Hexagonal Architecture

## Request Implement

user
  -POST   /register
  -POST   /login
  -PUT    /user/edit
  -DELETE /user/delete

book
  -GET    /books
  -GET    /books/:name
  -PUT    /books/:id
  -DELETE /books/:id

## About

DB: postgresql
ORM: gorm
HTTP: fiber
Tool: docker

### HOW TO START

docker-compose up -d
