package main

import (
	"fmt"
	"github.com/jochasinga/requests"
)

/*
var endpoints = {
  ticker: new UrlPattern('/ticker(?api_code=:apiCode)'),
  frombtc: new UrlPattern('/frombtc?value=:value&time=:time&currency=:currency(&api_code=:apiCode)'),
  tobtc: new UrlPattern('/tobtc?value=:value&currency=:currency(&api_code=:apiCode)')
}
*/

type API struct {
	baseUrl   string
	endpoints map[string]string
}

var exchangeEndpoints = map[string]string{
	"ticker":  "/ticker",
	"frombtc": "/frombtc",
	"tobtc":   "/tobtc",
}

func (api API) Get() string {
	return api.baseUrl
}

func main() {
	api := API{
		baseUrl:   "https://blockchain.info",
		endpoints: exchangeEndpoints,
	}

	jsonType := func(r *requests.Request) {
		r.Header.Add("content-type", "application/json")
	}

	res, err := requests.Get(api.baseUrl+api.endpoints["ticker"], jsonType)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", res.String())
}
