package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BinanceCryptoResponseRate struct {
	Symbol    string `json:"symbol"`
	LastPrice string `json:"lastPrice"`
	AskPrice  string `json:"askPrice"`
	BidPrice  string `json:"bidPrice"`
}

type BinanceCryptoResponseRates []BinanceCryptoResponseRate

func (rates BinanceCryptoResponseRates) transformRate() []RefinedRate {
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

func BinanceGetRates() ([]byte, error) {
	rates, err := BinanceFetchRates()

	if err != nil {
		return []byte{}, err
	}

	transformedRates := rates.transformRate()

	fmt.Println("Transformed rates", transformedRates)

	rateByte, err := json.Marshal(transformedRates)

	if err != nil {
		return []byte{}, err
	}

	return rateByte, nil
}

func BinanceFetchRates() (BinanceCryptoResponseRates, error) {
	var rates BinanceCryptoResponseRates
	response, error := http.Get("https://api.binance.com/api/v3/ticker/24hr")

	if error != nil {
		return BinanceCryptoResponseRates{}, error
	}

	byte, error := io.ReadAll(response.Body)

	if error != nil {
		return BinanceCryptoResponseRates{}, error
	}

	json.Unmarshal(byte, &rates)

	fmt.Println("Rate response", rates)

	return rates, nil
}
