package worker

import (
	"bpt/exchange"
	log "github.com/Sirupsen/logrus"
	"time"
)

func New(p exchange.Point, stop chan bool, results chan<- exchange.TradeData) {
	t := time.NewTicker(time.Second * p.Lifetime)

	go func() {
		data, err := p.Fetch()
		if err != nil {
			log.Warn(err.Error())
		} else {
			results <- data
		}

		for {
			select {
			case <-stop:
				break

			case <-t.C:
				data, err := p.Fetch()
				if err != nil {
					log.Warn(err.Error())
				} else {
					results <- data
				}
			}
		}
	}()
}
