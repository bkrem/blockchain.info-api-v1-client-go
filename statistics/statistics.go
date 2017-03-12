package statistics

import "github.com/bkrem/blockchain.info-api-v1-client-go/api"

type statisticsOpts struct {
	api.Opts
	StatName string `url:"stat"`
	Timespan string `url:"timespan,omitempty"`
	Format   string `url:"format"`
}

var endpoints = map[string]string{
	"charts": "/charts/",
	"pools":  "/pools?",
	"stats":  "/stats?",
}

var client = api.API{BaseURL: "https://blockchain.info", Endpoints: endpoints}

// GetChartData fetches the graphing data used in https://blockchain.info/charts
// for the specified `chartType`.
// `timespan` (e.g. "5weeks"); defaults to "1year" if passed an empty string.
// `format` can be "json" or 'csv'; defaults to to "json" if passed an empty string.
func GetChartData(chartType string, timespan string, format string) (string, error) {
	statsOpts := statisticsOpts{
		Opts:     api.Opts{},
		Timespan: timespan,
		Format:   format,
	}
	opts := chartType + "?" + client.EncodeOpts(statsOpts)
	res, err := client.Get("charts", opts)
	return res, err
}

// GetPoolData fetches data behind https://blockchain.info/pools.
// `timespan` has a maximum of "10days"; defaults to "4days" if passed an empty string.
func GetPoolData(timespan string) (string, error) {
	statsOpts := statisticsOpts{
		Opts:     api.Opts{},
		Timespan: timespan,
		Format:   "json",
	}
	opts := client.EncodeOpts(statsOpts)
	res, err := client.Get("pools", opts)
	return res, err
}

// GetStats fetches the data behind https://blockchain.info/stats.
func GetStats() (string, error) {
	statsOpts := statisticsOpts{
		Opts:   api.Opts{},
		Format: "json",
	}
	opts := client.EncodeOpts(statsOpts)
	res, err := client.Get("stats", opts)
	return res, err
}
