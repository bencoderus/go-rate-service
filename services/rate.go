package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bencoderus/go-rate-service/database/redis"
	"github.com/bencoderus/go-rate-service/utils"
)

const RATE_CACHE_KEY = "rateCacheKey"
const RATE_VALIDITY_IN_MINUTES = 3
const DURATION_IN_SECONDS = 60_000_000_000

type ConvertPayload struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

type RefinedRate struct {
	Currency  string `json:"currency"`
	Price     string `json:"price"`
	BuyPrice  string `json:"buyPrice"`
	SellPrice string `json:"sellPrice"`
}

func ConvertRateFromByteToStruct(rateBytes []byte) ([]RefinedRate, error) {
	var rates []RefinedRate
	err := json.Unmarshal(rateBytes, &rates)

	return rates, err
}

func GetRates() ([]byte, error) {
	redisClient := redis.GetRedisConnection()

	cached, error := redisClient.Get(RATE_CACHE_KEY)

	if error != nil {
		return nil, error
	}

	if cached != nil {
		return cached, nil
	}

	rateBytes, error := BinanceGetRates()

	if error != nil {
		return nil, error
	}

	redisClient.Set(RATE_CACHE_KEY, rateBytes, time.Duration(RATE_VALIDITY_IN_MINUTES*DURATION_IN_SECONDS))

	return rateBytes, nil
}

func validateConvertRatePayload(payload ConvertPayload) bool {
	return payload.Amount != 0 && payload.From != "" && payload.To != ""
}

func ConvertRate(payload []byte) utils.Error {
	var parsedPayload ConvertPayload
	json.Unmarshal(payload, &parsedPayload)

	if !validateConvertRatePayload(parsedPayload) {
		return utils.ThrowValidationError(fmt.Errorf("invalid payload"))
	}
	fmt.Println(parsedPayload)
	return utils.Error{}
}
