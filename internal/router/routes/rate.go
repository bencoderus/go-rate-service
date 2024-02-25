package routes

import (
	"github.com/bencoderus/go-rate-service/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func AddRateRoutes(app *fiber.App) {
	app.Get("/rates", handlers.GetRates)
	app.Post("/rates/convert", handlers.ConvertRate)
}
