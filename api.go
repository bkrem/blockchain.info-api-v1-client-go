package main

import (
	"fmt"
	"github.com/bkrem/blockchain.info-api-v1-client-go/exchange"
	"github.com/google/go-querystring/query"
	"github.com/jochasinga/requests"
)

type API struct {
	BaseURL   string
	Endpoints map[string]string
}

func (api API) Get() string {
	return api.BaseURL
}

func main() {
	api := API{"https://blockchain.info", exchange.Endpoints}

	jsonType := func(r *requests.Request) {
		r.Header.Add("content-type", "application/json")
	}

	opts := exchange.Options{Currency: "USD", Value: 500}
	v, _ := query.Values(opts)
	url := api.BaseURL + api.Endpoints["tobtc"] + v.Encode()

	res, err := requests.Get(url, jsonType)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query: %s\n", v.Encode())
	fmt.Printf("Full URL: %s\n", url)
	fmt.Print(res.String())
}
