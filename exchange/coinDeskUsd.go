package exchange

import (
	"encoding/json"
	"time"
)

type CoinDeskUsd struct{}

type CoinDeskUsdResponse struct {
	Bpi struct {
		USD struct {
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
	} `json:"bpi"`
}

func init() {
	RegisterParser("coinDeskUsd", CoinDeskUsd{})
}

func (coinDeskUsd CoinDeskUsd) Parse(p Point, raw json.RawMessage) (tD TradeData, err error) {
	tD.Name = p.Name
	tD.Pair = p.Pair

	var r CoinDeskUsdResponse
	err = json.Unmarshal(raw, &r)
	if err != nil {
		return
	}

	tD.Price = r.Bpi.USD.RateFloat
	tD.Time = time.Now()
	tD.ExpiredAt = tD.Time.Add(p.Lifetime * time.Second)

	return
}