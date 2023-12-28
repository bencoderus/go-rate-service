package handlers

import (
	"github.com/bencoderus/go-rate-service/services"
	"github.com/bencoderus/go-rate-service/utils"
	"github.com/gofiber/fiber/v2"
)

func GetRates(c *fiber.Ctx) error {
	rates, error := services.GetRates()

	if error != nil {
		response := utils.BuildJsonResponse(503, "Unable to get rates.")

		return c.Status(503).JSON(response)
	}

	rateResponse, _ := services.ConvertRateFromByteToStruct(rates)

	return c.Status(200).JSON(utils.BuildJsonResponseWithData(200, "Rate retrieved successfully.", rateResponse))
}

func ConvertRate(c *fiber.Ctx) error {
	services.ConvertRate(c.Body())
	response := utils.BuildJsonResponse(200, "Rate converted successfully.")

	return c.Status(200).JSON(response)
}
