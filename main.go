package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hamwiwatsapon/go-crud-authen/adapters"
	core "github.com/hamwiwatsapon/go-crud-authen/core"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	// dsn string connect db
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	// set logging for check connection
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	// User Secondary Port
	userRepo := adapters.NewGormUser(db)
	bookRepo := adapters.NewGormBook(db)

	// User Primary Port
	userService := core.NewUserService(userRepo)
	bookService := core.NewBookService(bookRepo)

	// User Primary Adapter
	userHandler := adapters.NewHttpUserHandler(userService)
	bookHandler := adapters.NewHttpBookHandler(bookService)

	// User Service
	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Authentication)
	app.Use("/user", authRequired)
	app.Put("/user/edit", userHandler.Edit)
	app.Delete("/user/delete", userHandler.Delete)

	// Book Service
	app.Use("/book", authRequired)
	app.Post("/books", bookHandler.NewBook)
	app.Get("/books", bookHandler.ReadBooks)
	app.Get("/books/:name", bookHandler.ReadNameBook)
	app.Put("/books/:id", bookHandler.UpdateBook)
	app.Delete("/books/:id", bookHandler.DeleteBook)
	app.Post("/books/author", bookHandler.NewAuthor)
	app.Post("/books/publisher", bookHandler.NewPublisher)

	db.AutoMigrate(&core.User{}, &core.Book{}, &core.Author{}, &core.AuthorBook{}, &core.Genre{}, &core.Publisher{})
	app.Listen(":8000")
}

func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	jwtSecretKey := os.Getenv("jwtSeCretKey")

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
