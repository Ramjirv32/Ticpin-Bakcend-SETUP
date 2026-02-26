package events

import (
	ctrl "ticpin-backend/controller/organizer/events"

	"github.com/gofiber/fiber/v2"
)

func EventsRoutes(app *fiber.App) {
	events := app.Group("/api/organizer/events")
	events.Post("/login", ctrl.EventsLogin)
	events.Post("/verify", ctrl.VerifyOTP)
	events.Post("/submit-verification", ctrl.SubmitVerification)
}
