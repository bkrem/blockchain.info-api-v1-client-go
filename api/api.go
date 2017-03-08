package api

import (
	"errors"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/jochasinga/requests"
)

type Opts struct {
	APICode string `url:"api_code,omitempty"`
}

type API struct {
	BaseURL   string
	Endpoints map[string]string
}

func (api API) Get(endpoint string, opts string) (string, error) {
	url := api.BaseURL + api.Endpoints[endpoint] + opts
	fmt.Printf("Get %s\n", url)
	res, err := requests.Get(url)
	if res.Response.StatusCode != 200 {
		return "", errors.New(res.String())
	}
	return res.String(), err
}

func (API) EncodeOpts(opts interface{}) string {
	v, err := query.Values(opts)
	if err != nil {
		panic(err)
	}
	encodedOpts := v.Encode()
	return encodedOpts
}
