package worker

import (
	"time"
	"bpt/exchange"
)

func New(p exchange.Point, stop chan bool, results chan<- exchange.TradeData) {
	t := time.NewTicker(time.Second * p.Lifetime)
	go func() {
		for {
			select {
			case <-stop:
				break

			case <-t.C:
				data, err := p.Fetch()
				if err != nil {
					//TODO: error handle
				}

				results <- data
			}
		}
	}()
}
