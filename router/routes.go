package router

import (
	"github.com/bencoderus/go-rate-service/handlers"
	"github.com/gofiber/fiber/v2"
)

func HandleRouting(app *fiber.App) {
	app.Get("/", handlers.HandleHome)
	app.Get("/rates", handlers.GetRates)
}
