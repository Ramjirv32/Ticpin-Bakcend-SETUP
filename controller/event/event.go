package event

import (
	"ticpin-backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllEvents(c *fiber.Ctx) error {
	events, err := services.GetAllEvents()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(events)
}

func GetEventByID(c *fiber.Ctx) error {
	e, err := services.GetEventByID(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "event not found"})
	}
	return c.JSON(e)
}
