package exchange

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"crypto/tls"
	"fmt"
	"time"
)

type Parser interface {
	Parse(p Point, raw json.RawMessage) (tD TradeData, err error)
}

type Point struct {
	HttpClient *http.Client
	Name      string
	Url       string
	Parser    Parser
	Lifetime  time.Duration
}

type TradeData struct {
	Name      string
	Price     float64
	Time      time.Time
	ExpiredAt time.Time
}

var (
	client *http.Client
	registeredParsers = map[string]Parser{}
)

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{Transport: tr}
}

func RegisterParser(name string, p Parser) {
	registeredParsers[name] = p
}

func NewPoint(name string, url string, parserName string, lifetime int) (p Point, err error) {
	parser, ok := registeredParsers[parserName]
	if !ok {
		err = fmt.Errorf("Not found parser with name: %s", parserName)
		return
	}

	p = Point{
		HttpClient: client,
		Name: name,
		Url:  url,
		Parser: parser,
		Lifetime: time.Duration(lifetime),
	}

	return
}

func (p *Point) Fetch() (tradeData TradeData, err error) {
	tradeData.Name = p.Name

	resp, err := p.HttpClient.Get(p.Url)
	if err != nil {
		err = fmt.Errorf("Error occurred fetching data from exchange [%s]: %s", p.Name, err.Error())
		return
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Error occurred reading data from exchange [%s]: %s", p.Name, err.Error())
		return
	}

	var data json.RawMessage
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		err = fmt.Errorf("Error occurred parsing data from exchange [%s]: %s", p.Name, err.Error())
		return
	}

	tradeData, err = p.Parser.Parse(*p, data)

	return
}
