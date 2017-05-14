package exchange

import (
	"encoding/json"
	"time"
)

type FixerUsd struct{}

type FixerUsdResponse struct {
	Rates struct {
		USD float64 `json:"USD"`
	} `json:"rates"`
}

func init() {
	RegisterParser("fixerUsd", FixerUsd{})
}

func (fixerUsd FixerUsd) Parse(p Point, raw json.RawMessage) (tD TradeData, err error) {
	tD.Name = p.Name
	tD.Pair = p.Pair

	var r FixerUsdResponse
	err = json.Unmarshal(raw, &r)
	if err != nil {
		return
	}

	tD.Price = r.Rates.USD
	tD.Time = time.Now()
	tD.ExpiredAt = tD.Time.Add(p.Lifetime * time.Second)

	return
}
