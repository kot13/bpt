package exchange

import (
	"encoding/json"
	"time"
)

type CoinDeskEur struct{}

type CoinDeskEurResponse struct {
	Bpi struct {
		EUR struct {
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func init() {
	RegisterParser("coinDeskEur", CoinDeskEur{})
}

func (coinDeskEur CoinDeskEur) Parse(p Point, raw json.RawMessage) (tD TradeData, err error) {
	tD.Name = p.Name
	tD.Pair = p.Pair

	var r CoinDeskEurResponse
	err = json.Unmarshal(raw, &r)
	if err != nil {
		return
	}

	tD.Price = r.Bpi.EUR.RateFloat
	tD.Time = time.Now()
	tD.ExpiredAt = tD.Time.Add(p.Lifetime * time.Second)

	return
}