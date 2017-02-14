package exchange

import (
	"fmt"
	"github.com/bkrem/blockchain.info-api-v1-client-go/api"
	"github.com/google/go-querystring/query"
	"regexp"
	"strconv"
)

type queryOpts struct {
	Currency string `url:"currency"`
	Value    string `url:"value"`
	Time     int    `url:"time"`
	APICode  string `url:"apiCode"`
}

var endpoints = map[string]string{
	"ticker":  "/ticker?",
	"frombtc": "/frombtc?",
	"tobtc":   "/tobtc?",
}

var client = api.API{BaseURL: "https://blockchain.info", Endpoints: endpoints}

func GetTicker() string {
	res := client.Get("ticker")
	fmt.Println(res)
	return res
}

// FIXME query string seems ineffectual?
/*
func GetTickerForCurrency(currency string) string {
	opts := encodeOpts(queryOpts{Currency: currency})
	res := client.GetWithOpts("ticker", opts)
	fmt.Println(res)
	return res
}
*/

func ToBTC(amount float64, currency string) (float64, error) {
	amountString := strconv.FormatFloat(amount, 'f', -1, 64)
	opts := encodeOpts(queryOpts{Currency: currency, Value: amountString})
	res := client.GetWithOpts("tobtc", opts)
	parsedRes, err := stringToFloat64(res)
	fmt.Println(parsedRes)
	return parsedRes, err
}

func FromBTC(amount int, currency string) (float64, error) {
	amountString := strconv.Itoa(amount)
	opts := encodeOpts(queryOpts{Currency: currency, Value: amountString})
	res := client.GetWithOpts("frombtc", opts)
	parsedRes, err := stringToFloat64(res)
	fmt.Println(parsedRes)
	return parsedRes, err
}

func encodeOpts(opts queryOpts) string {
	v, _ := query.Values(opts)
	encodedOpts := v.Encode()
	return encodedOpts
}

func stringToFloat64(floatString string) (float64, error) {
	// Remove comma separators in response to parse as float
	formattedRes := regexp.MustCompile(",").ReplaceAllString(floatString, "")
	// Cast the formatted string to float64
	parsedRes, err := strconv.ParseFloat(formattedRes, 64)
	return parsedRes, err
}
