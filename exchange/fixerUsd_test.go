package exchange

import (
	"encoding/json"
	"testing"
)

func TestFixerUsdParse(t *testing.T) {
	var raw json.RawMessage = json.RawMessage(`
		{"base": "EUR","date": "2017-05-12","rates": {"AUD": 1.4731,"BGN": 1.9558,"BRL": 3.4227,"CAD": 1.4941,"CHF": 1.0963,"CNY": 7.5047,"CZK": 26.576,"DKK": 7.4402,"GBP": 0.84588,"HKD": 8.4761,"HRK": 7.4225,"HUF": 310.24,"IDR": 14497.0,"ILS": 3.9203,"INR": 69.94,"JPY": 123.82,"KRW": 1226.5,"MXN": 20.552,"MYR": 4.7243,"NOK": 9.3665,"NZD": 1.5892,"PHP": 54.087,"PLN": 4.217,"RON": 4.545,"RUB": 62.315,"SEK": 9.6673,"SGD": 1.5314,"THB": 37.778,"TRY": 3.9038,"USD": 1.0876,"ZAR": 14.634}}
	`)

	parser := FixerUsd{}
	RegisterParser("test", parser)
	point, err := NewPoint("test", "test", "test", "test", 1)
	if err != nil {
		t.Error(err)
	}

	tradeData, err := parser.Parse(point, raw)
	if err != nil {
		t.Error(err)
	}

	if tradeData.Price != 1.0876 {
		t.Error("Error parse in FixerUsd parser")
	}
}
