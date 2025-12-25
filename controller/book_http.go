package controller

import (
	"go-clean-arc/entities"
	"go-clean-arc/usecases"

	"github.com/gofiber/fiber/v2"
)

type BookHttpHandler struct {
	usecase usecases.BookUseCase
}

func NewBookHttpHandler(usecase usecases.BookUseCase) *BookHttpHandler {
	return &BookHttpHandler{usecase: usecase}
}

func (u *BookHttpHandler) CreateBook(c *fiber.Ctx) error {
	var book entities.Book

	if err := c.BodyParser(&book); err != nil {
		return err
	}
	u.usecase.CreateBook(book)

	return nil
}

func (u *BookHttpHandler) FirstBook(c *fiber.Ctx) error {

	book := u.usecase.FirstBook()

	return c.JSON(book)
}
