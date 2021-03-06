package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

var cfg Cfg

type Cfg struct {
	App    AppConfig    `toml:"app"`
	Logger LoggerConfig `toml:"logger"`
	Feed   []FeedConfig `toml:"feed"`
}

type AppConfig struct {
	Interval int `toml:"interval"`
}

type LoggerConfig struct {
	LogLevel string `toml:"log_level"`
	LogFile  string `toml:"log_file"`
}

type FeedConfig struct {
	Name     string `toml:"name"`
	Pair     string `toml:"pair"`
	Url      string `toml:"url"`
	Parser   string `toml:"parser"`
	Lifetime int    `toml:"lifetime"`
}

func init() {
	fileName := flag.String("c", "config.toml", "config file name")

	flag.Parse()
	_, err := toml.DecodeFile(*fileName, &cfg)
	if err != nil {
		log.Fatal("decode: ", err)
		return
	}
}

func GetConfig() Cfg {
	return cfg
}
