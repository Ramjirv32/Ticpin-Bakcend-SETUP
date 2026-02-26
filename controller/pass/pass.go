package pass

import (
	"ticpin-backend/models"
	"ticpin-backend/services"

	"github.com/gofiber/fiber/v2"
)

func ApplyPass(c *fiber.Ctx) error {
	var req struct {
		UserID    string `json:"user_id"`
		PaymentID string `json:"payment_id"`
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
		Country   string `json:"country"`
		State     string `json:"state"`
		District  string `json:"district"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if req.UserID == "" || req.PaymentID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "user_id and payment_id required"})
	}

	details := models.TicpinPass{
		Name:     req.Name,
		Phone:    req.Phone,
		Address:  req.Address,
		Country:  req.Country,
		State:    req.State,
		District: req.District,
	}

	pass, err := services.ApplyPass(req.UserID, req.PaymentID, details)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(pass)
}

func GetPassByUser(c *fiber.Ctx) error {
	pass, err := services.GetActivePassByUserID(c.Params("userId"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "no active pass found"})
	}
	return c.JSON(pass)
}

func RenewPass(c *fiber.Ctx) error {
	var req struct {
		PaymentID string `json:"payment_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if req.PaymentID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "payment_id required"})
	}

	pass, err := services.RenewPass(c.Params("id"), req.PaymentID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pass)
}

func UseTurfBooking(c *fiber.Ctx) error {
	pass, err := services.UseTurfBooking(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pass)
}

func UseDiningVoucher(c *fiber.Ctx) error {
	pass, err := services.UseDiningVoucher(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pass)
}
