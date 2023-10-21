package handlers

import (
	"github.com/bencoderus/go-rate-service/services"
	"github.com/bencoderus/go-rate-service/utils"
	"github.com/gofiber/fiber/v2"
)

func GetRates(c *fiber.Ctx) error {
	rates, error := services.FetchRatesFromBinance()

	if error != nil {
		response := utils.BuildJsonResponse(503, "Unable to get rates.")

		return c.Status(503).JSON(response)
	}

	response := utils.BuildJsonResponseWithData(200, "Rate retrieved successfully.", rates)

	return c.Status(200).JSON(response)
}
