package handlers

import (
	"fmt"

	"github.com/bencoderus/go-rate-service/internal/services"
	"github.com/bencoderus/go-rate-service/pkg/utils"
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

func validatePayload(payload services.ConvertPayload) (bool, []string) {
	var errors []string
	if payload.To == "" {
		errors = append(errors, "to field is required.")
	}

	if payload.From == "" {
		errors = append(errors, "from field is required.")
	}

	if payload.Amount == 0 {
		errors = append(errors, "amount field is required.")
	}

	if payload.From != "" && services.SUPPORTED_COINS[payload.From] == "" {
		errors = append(errors, fmt.Sprintf("%s is not supported.", payload.From))
	}

	if payload.To != "" && services.SUPPORTED_COINS[payload.To] == "" {
		errors = append(errors, fmt.Sprintf("%s is not supported.", payload.To))
	}

	return len(errors) == 0, errors
}

func ConvertRate(c *fiber.Ctx) error {
	var payload services.ConvertPayload
	c.BodyParser(&payload)

	valid, errors := validatePayload(payload)

	if !valid {
		response := utils.BuildJsonResponseForValidationError(errors)

		return c.Status(200).JSON(response)
	}

	services.ConvertRate(payload)
	response := utils.BuildJsonResponse(200, "Rate converted successfully.")

	return c.Status(200).JSON(response)
}
