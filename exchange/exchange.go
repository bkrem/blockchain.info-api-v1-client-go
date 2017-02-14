package exchange

import (
	"fmt"
	"github.com/bkrem/blockchain.info-api-v1-client-go/api"
	"github.com/bkrem/blockchain.info-api-v1-client-go/util"
)

type exchangeOpts struct {
	api.Opts
	Currency string `url:"currency"`
	Value    string `url:"value"`
	Time     int    `url:"time,omitempty"`
}

var endpoints = map[string]string{
	"ticker":  "/ticker?",
	"frombtc": "/frombtc?",
	"tobtc":   "/tobtc?",
}

var client = api.API{BaseURL: "https://blockchain.info", Endpoints: endpoints}

func GetTicker() (string, error) {
	res, err := client.Get("ticker")
	return res, err
}

// FIXME query string seems ineffectual (#1)
/*
func GetTickerForCurrency(currency string) string {
	opts := encodeOpts(exchangeOpts{Currency: currency})
	res := client.GetWithOpts("ticker", opts)
	fmt.Println(res)
	return res
}
*/

func ToBTC(amount float64, currency string) (float64, error) {
	amountStr := util.Float64ToString(amount)
	eo := exchangeOpts{Opts: api.Opts{}, Currency: currency, Value: amountStr}
	opts := client.EncodeOpts(eo)
	fmt.Printf("OPTS: %s\n\n\nEND", opts)
	res, err := client.GetWithOpts("tobtc", opts)
	if err != nil {
		return 0, err
	}
	parsedRes, err := util.StringToFloat64(res)
	fmt.Println(parsedRes)
	return parsedRes, err
}

func FromBTC(amount int, currency string) (float64, error) {
	amountStr := util.IntToString(amount)
	eo := exchangeOpts{Opts: api.Opts{}, Currency: currency, Value: amountStr}
	opts := client.EncodeOpts(eo)
	res, err := client.GetWithOpts("frombtc", opts)
	if err != nil {
		return 0, err
	}
	parsedRes, err := util.StringToFloat64(res)
	fmt.Println(parsedRes)
	return parsedRes, err
}
