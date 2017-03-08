package statistics

import "github.com/bkrem/blockchain.info-api-v1-client-go/api"

/*
charts: new UrlPattern('/charts/:type?format=json(&api_code=:apiCode)(&timespan=:timespan)'),
pools: new UrlPattern('/pools?format=json(&timespan=:timespan\\days)(&api_code=:apiCode)'),
stats: new UrlPattern('/stats?format=json(&api_code=:apiCode)')
*/

type statisticsOpts struct {
	api.Opts
	Timespan string `url:"timespan,omitempty"`
	Format   string `url:"format"`
}

var endpoints = map[string]string{
	"charts": "/charts/",
	"pools":  "/pools?",
	"stats":  "/stats?",
}

var client = api.API{BaseURL: "https://blockchain.info", Endpoints: endpoints}

func GetChartData(chartType string, timespan string) (string, error) {
	statsOpts := statisticsOpts{
		Opts:     api.Opts{},
		Timespan: timespan,
		Format:   "json",
	}
	opts := chartType + "?" + client.EncodeOpts(statsOpts)
	res, err := client.Get("charts", opts)
	return res, err
}
