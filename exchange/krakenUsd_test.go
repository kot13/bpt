package exchange

import (
	"encoding/json"
	"testing"
)

func TestKrakenUsdParse(t *testing.T) {
	var raw json.RawMessage = json.RawMessage(`
		{"error": [ ],"result": {"XXBTZUSD": {"a": ["1772.99000","1","1.000"],"b": ["1769.85400","1","1.000"],"c": ["1773.83200","0.13796180"],"v": ["2146.99964507","2852.59095456"],"p": ["1788.13388","1782.45618"],"t": [4374,5878],"l": ["1752.85000","1723.16000"],"h": ["1818.00000","1818.00000"],"o": "1783.57000"}}}
	`)

	parser := KrakenUsd{}
	RegisterParser("test", parser)
	point, err := NewPoint("test", "test", "test", "test", 1)
	if err != nil {
		t.Error(err)
	}

	tradeData, err := parser.Parse(point, raw)
	if err != nil {
		t.Error(err)
	}

	if tradeData.Price != 1773.83200 {
		t.Error("Error parse in KrakenUsd parser")
	}
}
