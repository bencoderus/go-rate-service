package handlers

import (
	"github.com/bencoderus/go-rate-service/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleHome(c *fiber.Ctx) error {
	response := utils.BuildJsonResponse(200, "Rate service is active")

	return c.JSON(response)
}
