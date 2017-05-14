package ticker

import (
	"bpt/exchange"
	"time"
)

type Writer interface {
	Write(storage Storage)
}

type Storage map[string]map[string]Rate

type Rate struct {
	Value     float64
	ExpiredAt time.Time
}

var t *time.Ticker
var s Storage = make(Storage)

func Subscribe(results <-chan exchange.TradeData) {
	go func() {
		for {
			select {
			case result := <-results:
				rate, ok := s[result.Pair]
				if !ok {
					rate = make(map[string]Rate)
					s[result.Pair] = rate
				}

				s[result.Pair][result.Name] = Rate{
					Value:     result.Price,
					ExpiredAt: result.ExpiredAt,
				}
			}
		}
	}()
}

func Start(interval int, w Writer) {
	t = time.NewTicker(time.Second * time.Duration(interval))
	for range t.C {
		w.Write(s)
	}
}

func Stop() {
	t.Stop()
}
