package exchange

type Options struct {
	Currency string `url:"currency"`
	Value    int    `url:"value"`
	Time     int    `url:"time"`
	APICode  string `url:"apiCode"`
}

var Endpoints = map[string]string{
	"ticker":  "/ticker",
	"frombtc": "/frombtc?",
	"tobtc":   "/tobtc?",
}
