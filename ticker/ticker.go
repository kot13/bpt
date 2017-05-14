package ticker

import (
	"bpt/exchange"
	"sync"
	"time"
)

type Writer interface {
	Write(storage Storage)
}

type Storage struct {
	sync.RWMutex
	m map[string]map[string]Rate
}

type Rate struct {
	Value     float64
	ExpiredAt time.Time
}

var t *time.Ticker
var s Storage = Storage{
	m: make(map[string]map[string]Rate),
}

func Subscribe(results <-chan exchange.TradeData) {
	go func() {
		for {
			select {
			case result := <-results:
				s.Lock()

				rate, ok := s.m[result.Pair]
				if !ok {
					rate = make(map[string]Rate)
					s.m[result.Pair] = rate
				}

				s.m[result.Pair][result.Name] = Rate{
					Value:     result.Price,
					ExpiredAt: result.ExpiredAt,
				}

				s.Unlock()
			}
		}
	}()
}

func Start(interval int, w Writer) {
	t = time.NewTicker(time.Second * time.Duration(interval))
	for range t.C {
		s.RLock()
		w.Write(s)
		s.RUnlock()
	}
}

func Stop() {
	t.Stop()
}
