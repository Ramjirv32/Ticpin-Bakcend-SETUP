package main

import (
	"log"
	"ticpin-backend/config"
	"ticpin-backend/routes/organizer"
	organizerplay "ticpin-backend/routes/organizer/play"
	"ticpin-backend/routes/profile"
	"ticpin-backend/routes/user"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := config.ConnectDB(); err != nil {
		log.Fatal("MongoDB:", err)
	}

	if err := config.InitCloudinary(); err != nil {
		log.Fatal("Cloudinary:", err)
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	user.UserRoutes(app)
	profile.ProfileRoutes(app)
	organizer.OrganizerRoutes(app)
	organizerplay.PlayRoutes(app)

	log.Fatal(app.Listen(":9000"))
}
