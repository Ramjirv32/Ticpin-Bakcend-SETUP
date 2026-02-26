package main

import (
	"log"
	"ticpin-backend/config"
	diningroutes "ticpin-backend/routes/dining"
	eventroutes "ticpin-backend/routes/event"
	"ticpin-backend/routes/organizer"
	organizerdining "ticpin-backend/routes/organizer/dining"
	organizerEvents "ticpin-backend/routes/organizer/events"
	organizerplay "ticpin-backend/routes/organizer/play"
	passroutes "ticpin-backend/routes/pass"
	playroutes "ticpin-backend/routes/play"
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
	organizerEvents.EventsRoutes(app)
	organizerdining.DiningRoutes(app)
	eventroutes.EventRoutes(app)
	playroutes.PlayRoutes(app)
	diningroutes.DiningRoutes(app)
	passroutes.PassRoutes(app)

	log.Fatal(app.Listen(":9000"))
}
