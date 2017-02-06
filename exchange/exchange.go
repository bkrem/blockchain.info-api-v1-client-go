package exchange

type Options struct {
	Currency string `url:"currency"`
	Value    int    `url:"value"`
	Time     int    `url:time`
	ApiCode  string `url:apiCode`
}

var Endpoints = map[string]string{
	"ticker":  "/ticker",
	"frombtc": "/frombtc?",
	"tobtc":   "/tobtc?",
}
