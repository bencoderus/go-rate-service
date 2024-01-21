package main

import (
	"errors"
	"fmt"
	"os"
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
				errorMessage = e.Message
			}

			if code >= 500 {
				fmt.Println("error", err)
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	fmt.Println("Server running on port", port)

	app.Listen(":3003")
}
