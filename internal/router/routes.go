package router

import (
	"github.com/bencoderus/go-rate-service/internal/handlers"
	"github.com/bencoderus/go-rate-service/internal/router/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func HandleRouting(app *fiber.App) {
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/", handlers.HandleHome)
	routes.AddRateRoutes(app)

	app.Get("/start", func(c *fiber.Ctx) error {
		panic("This panic is caught by fiber")
	})
}
