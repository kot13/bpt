package exchange

import (
	"encoding/json"
	"testing"
)

func TestCoinDeskUsdParse(t *testing.T) {
	var raw json.RawMessage = json.RawMessage(`
		{"time":{"updated":"May 14, 2017 19:02:00 UTC","updatedISO":"2017-05-14T19:02:00+00:00","updateduk":"May 14, 2017 at 20:02 BST"},"disclaimer":"This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org","chartName":"Bitcoin","bpi":{"USD":{"code":"USD","symbol":"&#36;","rate":"1,750.9400","description":"United States Dollar","rate_float":1750.94},"GBP":{"code":"GBP","symbol":"&pound;","rate":"1,358.4843","description":"British Pound Sterling","rate_float":1358.4843},"EUR":{"code":"EUR","symbol":"&euro;","rate":"1,601.4447","description":"Euro","rate_float":1601.4447}}}
	`)

	parser := CoinDeskUsd{}
	RegisterParser("test", parser)
	point, err := NewPoint("test", "test", "test", "test", 1)
	if err != nil {
		t.Error(err)
	}

	tradeData, err := parser.Parse(point, raw)
	if err != nil {
		t.Error(err)
	}

	if tradeData.Price != 1750.94 {
		t.Error("Error parse in CoinDeskUsd parser")
	}
}
