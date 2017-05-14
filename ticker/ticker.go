package ticker

import (
	"time"
	"bpt/exchange"
)

type Writer interface {
	Write(storage Storage)
}

type Storage map[string]map[string]Rate

type Rate struct {
	Value 		float64
	ExpiredAt	time.Time
}

var t *time.Ticker
var s Storage = make(Storage)

func Subscribe(results <-chan exchange.TradeData) {
	go func() {
		for {
			select {
			case result := <-results:
				go func() {
					rate, ok := s[result.Pair]
					if !ok {
						rate = make(map[string]Rate)
						s[result.Pair] = rate
					}

					s[result.Pair][result.Name] = Rate{
						Value: result.Price,
						ExpiredAt: result.ExpiredAt,
					}
				}()
			}
		}
	}()
}

func Start(interval int, w Writer) {
	t = time.NewTicker(time.Second * time.Duration(interval))
	go func() {
		for range t.C {
			w.Write(s)
		}
	}()
}

func Stop() {
	t.Stop()
}