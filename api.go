package main

import (
	"fmt"
	"github.com/bkrem/blockchain.info-api-v1-client-go/exchange"
	"github.com/google/go-querystring/query"
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
	BaseUrl   string
	Endpoints map[string]string
}

func (api API) Get() string {
	return api.BaseUrl
}

func main() {
	api := API{"https://blockchain.info", exchange.ExchangeEndpoints}

	jsonType := func(r *requests.Request) {
		r.Header.Add("content-type", "application/json")
	}

	opts := exchange.Options{"USD", 500}
	v, _ := query.Values(opts)
	url := api.BaseUrl + api.Endpoints["tobtc"] + v.Encode()

	res, err := requests.Get(url, jsonType)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query: %s\n", v.Encode())
	fmt.Printf("Full URL: %s\n", url)
	fmt.Print(res.String())
}
