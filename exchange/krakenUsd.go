package exchange

import (
	"encoding/json"
	"strconv"
	"time"
)

type KrakenUsd struct{}

type KrakenUsdResponse struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XXBTZUSD struct {
			C []string `json:"c"`
		} `json:"XXBTZUSD"`
	} `json:"result"`
}

func init() {
	RegisterParser("krakenUsd", KrakenUsd{})
}

func (krakenUsd KrakenUsd) Parse(p Point, raw json.RawMessage) (tD TradeData, err error) {
	tD.Name = p.Name
	tD.Pair = p.Pair

	var r KrakenUsdResponse
	err = json.Unmarshal(raw, &r)
	if err != nil {
		return
	}

	tD.Price, err = strconv.ParseFloat(r.Result.XXBTZUSD.C[0], 64)
	if err != nil {
		return
	}

	tD.Time = time.Now()
	tD.ExpiredAt = tD.Time.Add(p.Lifetime * time.Second)

	return
}