package blockexplorer

import (
	"github.com/bkrem/blockchain.info-api-v1-client-go/api"
	"github.com/bkrem/blockchain.info-api-v1-client-go/util"
)

/*
module.exports = {
  rawblock: new UrlPattern('/rawblock/:hash(?api_code=:apiCode)'),
  rawtx: new UrlPattern('/rawtx/:hash(?api_code=:apiCode)'),
  blockHeight: new UrlPattern('/block-height/:height?format=json(&api_code=:apiCode)'),
  address: new UrlPattern('/address/:address?format=json(&limit=:limit)(&offset=:offset)(&api_code=:apiCode)'),
  multiaddr: new UrlPattern('/multiaddr?active=:active(&n=:limit)(&offset=:offset)(&api_code=:apiCode)'),
  unspent: new UrlPattern('/unspent?active=:active(&api_code=:apiCode)'),
  latestblock: new UrlPattern('/latestblock(?api_code=:apiCode)'),
  unconfTxs: new UrlPattern('/unconfirmed-transactions?format=json(&api_code=:apiCode)'),
  blocks: new UrlPattern('/blocks/:time?format=json(&api_code=:apiCode)'),
  inv: new UrlPattern('/inv/:hash?format=json(&api_code=:apiCode)')
}
*/

type blockExplorerOpts struct {
	api.Opts
	Format string `url:"format,omitempty"`
}

var endpoints = map[string]string{
	"rawBlock":    "/rawblock/",
	"rawTx":       "/rawtx/",
	"blockHeight": "/block-height/",
	"address":     "/address/",
	"multiAddr":   "/multiaddr?",
	"unspent":     "/unspent?",
	"latestBlock": "/latestblock?",
	"unconfTxs":   "/unconfirmed-transactions?",
	"blocks":      "/blocks/",
	"inv":         "/inv/",
}

var client = api.API{BaseURL: "https://blockchain.info", Endpoints: endpoints}

func GetBlockByHash(hash string) (string, error) {
	beOpts := blockExplorerOpts{Opts: api.Opts{}}
	opts := hash + "?" + client.EncodeOpts(beOpts)
	res, err := client.GetWithOpts("rawBlock", opts)
	return res, err
}

func GetBlockByIndex(index int) (string, error) {
	beOpts := blockExplorerOpts{Opts: api.Opts{}}
	opts := util.IntToString(index) + "?" + client.EncodeOpts(beOpts)
	res, err := client.GetWithOpts("rawBlock", opts)
	return res, err
}

func GetBlockHeight(height int) (string, error) {
	beOpts := blockExplorerOpts{Opts: api.Opts{}, Format: "json"}
	opts := util.IntToString(height) + "?" + client.EncodeOpts(beOpts)
	res, err := client.GetWithOpts("blockHeight", opts)
	return res, err
}

func GetLatestBlock() (string, error) {
	res, err := client.Get("latestBlock")
	return res, err
}
