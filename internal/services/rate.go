package services

import (
	"encoding/json"
	"time"

	"github.com/bencoderus/go-rate-service/internal/database/redis"
	"github.com/bencoderus/go-rate-service/pkg/utils"
)

const RATE_CACHE_KEY = "rateCacheKey"
const RATE_VALIDITY_IN_MINUTES = 3
const DURATION_IN_SECONDS = 60_000_000_000

var SUPPORTED_COINS = map[string]string{
	"BTC":  "Bitcoin",
	"ETH":  "Ethereum",
	"SOL":  "Solana",
	"DOGE": "Doge",
	"USDT": "Tether",
}

type ConvertPayload struct {
	From   string `json:"from" validate:"required,min=3"`
	To     string `json:"to" validate:"required,min=3"`
	Amount int    `json:"amount" validate:"required,min=1"`
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
	redisClient, error := redis.GetRedisConnection()

	if error != nil {
		return BinanceGetRates()
	}

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

func ConvertRate(payload ConvertPayload) utils.Error {
	return utils.Error{}
}
