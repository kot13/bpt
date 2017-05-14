package ticker

import (
	"time"
	"fmt"
	"bpt/exchange"
	"log"
)

var t *time.Ticker

func Subscribe(results <-chan exchange.TradeData) {
	go func() {
		for {
			select {
			case result := <-results:
				go func() {
					// TODO: save to storage
					log.Print(result)
				}()
			}
		}
	}()
}

func Start(interval int) {
	t = time.NewTicker(time.Second * time.Duration(interval))
	go func() {
		for tick := range t.C {
			// TODO: Load data and render output line
			fmt.Println("Tick at ", tick)
		}
	}()
}

func Stop() {
	t.Stop()
}