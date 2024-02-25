package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bencoderus/go-rate-service/internal/router"
	"github.com/bencoderus/go-rate-service/pkg/utils"
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
				log.Println(err)
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

	fmt.Println("App running on port", port)

	app.Listen(fmt.Sprintf(":%s", port))
}
