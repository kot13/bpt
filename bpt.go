package main

import (
	log "github.com/Sirupsen/logrus"

	"bpt/config"
	"bpt/exchange"
	"bpt/logger"
	"bpt/ticker"
	"bpt/worker"
)

func main() {
	conf := config.GetConfig()
	logFinalizer, err := logger.InitLogger(conf.Logger.LogLevel, conf.Logger.LogFile)
	if err != nil {
		log.Fatal(err)
	}
	defer logFinalizer()

	log.Info("Start bpt")

	stop := make(chan bool)
	results := make(chan exchange.TradeData)
	for _, feed := range conf.Feed {
		point, err := exchange.NewPoint(feed.Name, feed.Pair, feed.Url, feed.Parser, feed.Lifetime)
		if err != nil {
			log.Fatal(err)
		}

		worker.New(point, stop, results)
	}

	ticker.Subscribe(results)
	ticker.Start(conf.App.Interval, ticker.SimpleWriter{})
}
