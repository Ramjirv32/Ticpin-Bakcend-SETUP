package organizer

import (
	ctrl "ticpin-backend/controller/organizer"

	"github.com/gofiber/fiber/v2"
)

func OrganizerRoutes(app *fiber.App) {
	profile := app.Group("/api/organizer/profile")
	profile.Post("", ctrl.CreateProfile)
	profile.Get("/:id", ctrl.GetProfile)
	profile.Put("/:id", ctrl.UpdateProfile)

	verification := app.Group("/api/organizer/verification")
	verification.Get("/:id", ctrl.GetVerificationStatus)
}
