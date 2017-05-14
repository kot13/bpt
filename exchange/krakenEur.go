package exchange

import (
	"encoding/json"
	"strconv"
	"time"
)

type KrakenEur struct{}

type KrakenEurResponse struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XXBTZEUR struct {
			C []string `json:"c"`
		} `json:"XXBTZEUR"`
	} `json:"result"`
}

func init() {
	RegisterParser("krakenEur", KrakenEur{})
}

func (krakenEur KrakenEur) Parse(p Point, raw json.RawMessage) (tD TradeData, err error) {
	tD.Name = p.Name
	tD.Pair = p.Pair

	var r KrakenEurResponse
	err = json.Unmarshal(raw, &r)
	if err != nil {
		return
	}

	tD.Price, err = strconv.ParseFloat(r.Result.XXBTZEUR.C[0], 64)
	if err != nil {
		return
	}

	tD.Time = time.Now()
	tD.ExpiredAt = tD.Time.Add(p.Lifetime * time.Second)

	return
}
