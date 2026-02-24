package profile

import (
	"ticpin-backend/models"
	"ticpin-backend/services"

	"github.com/gofiber/fiber/v2"
)

func CreateProfile(c *fiber.Ctx) error {
	var p models.Profile
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.CreateProfile(&p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(p)
}

func GetProfile(c *fiber.Ctx) error {
	p, err := services.GetProfileByUserID(c.Params("userId"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "profile not found"})
	}

	return c.JSON(p)
}

func UpdateProfile(c *fiber.Ctx) error {
	var p models.Profile
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.UpdateProfile(c.Params("userId"), &p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "profile updated"})
}

func UploadProfilePhoto(c *fiber.Ctx) error {
	userID := c.Params("userId")

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "photo required"})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer src.Close()

	photoURL, err := services.UploadProfilePhoto(src, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "upload failed"})
	}

	if err := services.UpdateProfilePhoto(userID, photoURL); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"photoURL": photoURL})
}
