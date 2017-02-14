package exchange

import (
	"fmt"
	"github.com/bkrem/blockchain.info-api-v1-client-go/api"
	"github.com/bkrem/blockchain.info-api-v1-client-go/util"
	"strconv"
)

type exchangeOpts struct {
	api.Opts
	Currency string `url:"currency"`
	Value    string `url:"value"`
	Time     int    `url:"time,omitempty"`
}

var endpoints = map[string]string{
	"ticker":  "/ticker",
	"frombtc": "/frombtc",
	"tobtc":   "/tobtc",
}

var client = api.API{BaseURL: "https://blockchain.info", Endpoints: endpoints}

func GetTicker() string {
	return client.Get("ticker")
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
	amountStr := strconv.FormatFloat(amount, 'f', -1, 64)
	eo := exchangeOpts{Opts: api.Opts{}, Currency: currency, Value: amountStr}
	opts := client.EncodeOpts(eo)
	fmt.Printf("OPTS: %s\n\n\nEND", opts)
	res := client.GetWithOpts("tobtc", opts)
	parsedRes, err := util.StringToFloat64(res)
	fmt.Println(parsedRes)
	return parsedRes, err
}

func FromBTC(amount int, currency string) (float64, error) {
	amountStr := strconv.Itoa(amount)
	eo := exchangeOpts{Opts: api.Opts{}, Currency: currency, Value: amountStr}
	opts := client.EncodeOpts(eo)
	res := client.GetWithOpts("frombtc", opts)
	parsedRes, err := util.StringToFloat64(res)
	fmt.Println(parsedRes)
	return parsedRes, err
}
