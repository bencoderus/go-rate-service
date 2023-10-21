package services

import (
	"encoding/json"
	"io"
	"net/http"
)

type BinanceRate struct {
	Symbol    string `json:"symbol"`
	LastPrice string `json:"lastPrice"`
	AskPrice  string `json:"askPrice"`
	BidPrice  string `json:"bidPrice"`
}

type RefinedRate struct {
	Currency  string `json:"currency"`
	Price     string `json:"price"`
	BuyPrice  string `json:"buyPrice"`
	SellPrice string `json:"sellPrice"`
}

func getBinanceCurrencyPairs() map[string]string {
	return map[string]string{
		"BTCUSDT":  "BTC",
		"ETHUSDT":  "ETH",
		"SOLUSDT":  "SOL",
		"XRPUSDT":  "XRP",
		"DOGEUSDT": "DOGE",
		"TUSDUSDT": "USDT",
	}
}

func TransformRates(rates []BinanceRate) []RefinedRate {
	var transformed []RefinedRate

	symbolCurrency := getBinanceCurrencyPairs()

	for _, rate := range rates {
		_, exists := symbolCurrency[rate.Symbol]

		if exists {
			transformed = append(transformed, RefinedRate{
				Currency:  symbolCurrency[rate.Symbol],
				BuyPrice:  rate.BidPrice,
				SellPrice: rate.AskPrice,
				Price:     rate.LastPrice,
			})
		}
	}

	return transformed
}

func FetchRatesFromBinance() ([]RefinedRate, error) {
	var rates []BinanceRate
	response, error := http.Get("https://api.binance.com/api/v3/ticker/24hr")

	if error != nil {
		return []RefinedRate{}, error
	}

	byte, error := io.ReadAll(response.Body)

	if error != nil {
		return []RefinedRate{}, error
	}

	json.Unmarshal(byte, &rates)

	transformed := TransformRates(rates)

	return transformed, nil

}
