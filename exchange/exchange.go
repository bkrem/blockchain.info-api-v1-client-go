package exchange

type Options struct {
	Currency string `url:"currency"`
	Value    int    `url:"value"`
}

var ExchangeEndpoints = map[string]string{
	"ticker":  "/ticker",
	"frombtc": "/frombtc?",
	"tobtc":   "/tobtc?",
}
