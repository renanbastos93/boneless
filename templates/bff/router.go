package bff

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/renanbastos93/boneless/templates/app"
)

func (e implBFF) GetAllExamples(c *fiber.Ctx) (err error) {
	out, err := e.example.Get().AllExamples(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) GetExampleById(c *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid parameter:"+c.Params("id"))
	}
	out, err := e.example.Get().GetOneExampleById(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) CreateExample(c *fiber.Ctx) (err error) {
	var body app.ExampleIn

	err = c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = e.example.Get().CreateExample(c.Context(), body)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusCreated)
}
