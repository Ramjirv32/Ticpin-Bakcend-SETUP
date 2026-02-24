package play

import (
	ctrl "ticpin-backend/controller/organizer/play"

	"github.com/gofiber/fiber/v2"
)

func PlayRoutes(app *fiber.App) {
	play := app.Group("/api/organizer/play")
	play.Post("/login", ctrl.PlayLogin)
	play.Post("/verify", ctrl.VerifyOTP)
	play.Get("/:id", ctrl.GetOrganizer)
}
