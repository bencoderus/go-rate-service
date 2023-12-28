package main

import (
	"errors"
	"fmt"

	"github.com/bencoderus/go-rate-service/router"
	"github.com/bencoderus/go-rate-service/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func CreateApp() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			errorMessage := "Unable to process request."

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if code >= 500 {
				fmt.Println("error", err)
			}

			if errors.As(err, &e) {
				errorMessage = e.Message
			}

			return ctx.Status(code).JSON(utils.BuildJsonResponse(code, errorMessage))
		},
	})
}

func main() {
	app := CreateApp()

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
	}
	router.HandleRouting(app)

	app.Listen(":3002")
}
