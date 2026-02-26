package play

import (
	"ticpin-backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllPlays(c *fiber.Ctx) error {
	plays, err := services.GetAllPlays()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(plays)
}

func GetPlayByID(c *fiber.Ctx) error {
	p, err := services.GetPlayByID(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "play not found"})
	}
	return c.JSON(p)
}
