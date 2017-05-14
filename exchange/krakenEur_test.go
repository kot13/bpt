package exchange

import (
	"encoding/json"
	"testing"
)

func TestKrakenEurParse(t *testing.T) {
	var raw json.RawMessage = json.RawMessage(`
		{"error": [ ],"result": {"XXBTZEUR": {"a": ["1624.32000","5","5.000"],"b": ["1620.18500","2","2.000"],"c": ["1624.32000","0.29491750"],"v": ["7484.64587730","11623.30264059"],"p": ["1640.00001","1634.76153"],"t": [13036,18950],"l": ["1608.46900","1590.55000"],"h": ["1668.00000","1668.00000"],"o": "1638.87000"}}}
	`)

	parser := KrakenEur{}
	RegisterParser("test", parser)
	point, err := NewPoint("test", "test", "test", "test", 1)
	if err != nil {
		t.Error(err)
	}

	tradeData, err := parser.Parse(point, raw)
	if err != nil {
		t.Error(err)
	}

	if tradeData.Price != 1624.32000 {
		t.Error("Error parse in KrakenEur parser")
	}
}
