package adapters

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	core "github.com/hamwiwatsapon/go-crud-authen/core"
)

// Primary Adapter receive from HTTP
type HttpBookHandler struct {
	service core.BookService
}

func NewHttpBookHandler(service core.BookService) *HttpBookHandler {
	return &HttpBookHandler{service: service}
}

func (h *HttpBookHandler) NewBook(c *fiber.Ctx) error {
	var book core.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.service.NewBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *HttpBookHandler) ReadBooks(c *fiber.Ctx) error {
	books, err := h.service.ReadBooks()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(books)
}

func (h *HttpBookHandler) ReadNameBook(c *fiber.Ctx) error {
	bookName := c.Params("name")

	books, err := h.service.ReadNameBook(bookName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(books)
}

func (h *HttpBookHandler) UpdateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var book core.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.service.UpdateBook(bookId, book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(book)
}

func (h *HttpBookHandler) DeleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.service.DeleteBook(bookId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusAccepted)
}
