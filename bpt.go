package main

import (
	"bpt/config"
	"bpt/logger"

	log "github.com/Sirupsen/logrus"
	"bpt/exchange"
)

func main()  {
	conf := config.GetConfig()
	logFinalizer, err := logger.InitLogger(conf.Logger.LogLevel, conf.Logger.LogFile)
	if err != nil {
		log.Fatal(err)
	}
	defer logFinalizer()

	log.Info("Start bpt")

	exchangePoints := []exchange.Point{}
	for _, feed := range conf.Feed {
		point, err := exchange.NewPoint(feed.Name, feed.Url, feed.Parser, feed.Lifetime)
		if err != nil {
			log.Fatal(err)
		}

		exchangePoints = append(exchangePoints, point)
	}

	data, err := exchangePoints[0].Fetch()
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v", data)

	data, err = exchangePoints[1].Fetch()
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v", data)
}