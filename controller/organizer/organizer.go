package organizer

import (
	"ticpin-backend/models"
	"ticpin-backend/services"

	"github.com/gofiber/fiber/v2"
)

func CreateProfile(c *fiber.Ctx) error {
	var p models.OrganizerProfile
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.CreateOrganizerProfile(&p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(p)
}

func GetProfile(c *fiber.Ctx) error {
	p, err := services.GetOrganizerProfileByID(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "profile not found"})
	}

	return c.JSON(p)
}

func UpdateProfile(c *fiber.Ctx) error {
	var p models.OrganizerProfile
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.UpdateOrganizerProfile(c.Params("id"), &p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "profile updated"})
}

func GetVerificationStatus(c *fiber.Ctx) error {
	v, err := services.GetOrganizerVerification(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "verification record not found"})
	}

	return c.JSON(v)
}
