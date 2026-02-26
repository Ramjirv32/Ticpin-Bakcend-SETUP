package dining

import (
	ctrl "ticpin-backend/controller/organizer/dining"

	"github.com/gofiber/fiber/v2"
)

func DiningRoutes(app *fiber.App) {
	dining := app.Group("/api/organizer/dining")
	dining.Post("/login", ctrl.DiningLogin)
	dining.Post("/verify", ctrl.VerifyOTP)
	dining.Post("/submit-verification", ctrl.SubmitVerification)
}
