package play

import (
	"ticpin-backend/services"

	"github.com/gofiber/fiber/v2"
)

func PlayLogin(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if req.Email == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "email and password required"})
	}

	org, isNew, err := services.LoginOrCreateOrganizer(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.SendPlayOrganizerOTP(req.Email); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to send otp"})
	}

	msg := "otp sent"
	if isNew {
		msg = "account created, otp sent"
	}

	return c.JSON(fiber.Map{"message": msg, "organizerId": org.ID})
}

func VerifyOTP(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	org, err := services.VerifyOrganizerOTP(req.Email, req.OTP)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(org)
}

func GetOrganizer(c *fiber.Ctx) error {
	org, err := services.GetOrganizerByID(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "organizer not found"})
	}

	return c.JSON(org)
}
