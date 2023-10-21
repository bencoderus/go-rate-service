package main

import (
	"github.com/bencoderus/go-rate-service/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.HandleRouting(app)

	app.Listen(":3002")
}
