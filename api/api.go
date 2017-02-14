package api

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/jochasinga/requests"
)

type Opts struct {
	APICode string `url:"api_code,omitempty"`
}

type OptsEncoder interface {
	EncodeOpts() string
}

type API struct {
	BaseURL   string
	Endpoints map[string]string
}

func (api API) Get(endpoint string) string {
	url := api.BaseURL + api.Endpoints[endpoint]
	fmt.Printf("Get %s\n", url)
	res, err := requests.Get(url)
	if err != nil {
		panic(err)
	}
	return res.String()
}

func (api API) GetWithOpts(endpoint string, opts string) string {
	url := api.BaseURL + api.Endpoints[endpoint] + "?" + opts
	fmt.Printf("GetWithOpts %s\n", url)
	res, err := requests.Get(url)
	if err != nil {
		panic(err)
	}
	return res.String()
}

func (API) EncodeOpts(opts interface{}) string {
	v, _ := query.Values(opts)
	encodedOpts := v.Encode()
	return encodedOpts
}
