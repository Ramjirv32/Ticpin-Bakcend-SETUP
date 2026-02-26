package dining

import (
	"ticpin-backend/models"
	"ticpin-backend/services"

	"github.com/gofiber/fiber/v2"
)

func DiningLogin(c *fiber.Ctx) error {
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

	if err := services.SendOrganizerOTP(req.Email, "dining"); err != nil {
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

func SubmitVerification(c *fiber.Ctx) error {
	var v models.DiningVerification
	if err := c.BodyParser(&v); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if v.OrganizerID.IsZero() {
		return c.Status(400).JSON(fiber.Map{"error": "organizer_id required"})
	}

	if err := services.SubmitDiningVerification(&v); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "verification submitted"})
}
