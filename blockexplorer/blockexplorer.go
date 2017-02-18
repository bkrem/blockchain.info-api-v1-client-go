package blockexplorer

import (
	"strings"

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
	Limit  int    `url:"limit,omitempty"`
	Active string `url:"active,omitempty"`
	Offset int    `url:"offset,omitempty"`
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

// GetBlockByHash fetches the block corresponding to the passed hash string
// Query pattern: `/rawblock/:hash(?api_code=:apiCode)`
func GetBlockByHash(hash string) (string, error) {
	beOpts := blockExplorerOpts{Opts: api.Opts{}}
	opts := hash + "?" + client.EncodeOpts(beOpts)
	res, err := client.Get("rawBlock", opts)
	return res, err
}

// GetBlockByIndex fetches the block corresponding to the passed index
// Query pattern: `/rawblock/:index(?api_code=:apiCode)`
func GetBlockByIndex(index int) (string, error) {
	beOpts := blockExplorerOpts{Opts: api.Opts{}}
	opts := util.IntToString(index) + "?" + client.EncodeOpts(beOpts)
	res, err := client.Get("rawBlock", opts)
	return res, err
}

// GetBlockHeight fetches the block corresponding to the passed block height
// Query pattern: `/block-height/:height?format=json(&api_code=:apiCode)`
func GetBlockHeight(height int) (string, error) {
	beOpts := blockExplorerOpts{
		Opts:   api.Opts{},
		Format: "json",
	}
	opts := util.IntToString(height) + "?" + client.EncodeOpts(beOpts)
	res, err := client.Get("blockHeight", opts)
	return res, err
}

// GetAddress fetches the passed bitcoin address's/hash160's transactions
// NOTE: Passing `0` for `limit` or `offset` will omit these options
// Query pattern: `/address/:address?format=json(&limit=:limit)(&offset=:offset)(&api_code=:apiCode)`
func GetAddress(address string, limit int, offset int) (string, error) {
	beOpts := blockExplorerOpts{
		Opts:   api.Opts{},
		Limit:  limit,
		Offset: offset,
		Format: "json",
	}
	opts := address + "?" + client.EncodeOpts(beOpts)
	res, err := client.Get("address", opts)
	return res, err
}

// GetMultiAddress fetches summaries of the passed addresses and their transactions
// NOTE: Passing `0` for `limit` or `offset` will omit these options
// Query pattern: `/multiaddr?active=:active(&n=:limit)(&offset=:offset)(&api_code=:apiCode)`
func GetMultiAddress(addresses []string, limit int, offset int) (string, error) {
	joinedAddresses := strings.Join(addresses, "|")
	beOpts := blockExplorerOpts{
		Opts:   api.Opts{},
		Active: joinedAddresses,
		Limit:  limit,
		Offset: offset,
	}
	opts := client.EncodeOpts(beOpts)
	res, err := client.Get("multiAddr", opts)
	return res, err
}

// GetUnspentOutputs fetches an array of unspent outputs for the passed addresses
// Query pattern: `/unspent?active=:active(&api_code=:apiCode)`
func GetUnspentOutputs(addresses []string) (string, error) {
	joinedAddresses := strings.Join(addresses, "|")
	beOpts := blockExplorerOpts{
		Opts:   api.Opts{},
		Active: joinedAddresses,
	}
	opts := client.EncodeOpts(beOpts)
	res, err := client.Get("unspent", opts)
	return res, err
}

// GetLatestBlock fetches the latest block on the main chain
// Query pattern: `/latestblock(?api_code=:apiCode)`
func GetLatestBlock() (string, error) {
	res, err := client.Get("latestBlock", "")
	return res, err
}

// GetUnconfirmedTxs fetches an array of all unconfirmed transactions
// Query pattern: `/unconfirmed-transactions?format=json(&api_code=:apiCode)`
func GetUnconfirmedTxs() (string, error) {
	beOpts := blockExplorerOpts{
		Opts:   api.Opts{},
		Format: "json",
	}
	opts := client.EncodeOpts(beOpts)
	res, err := client.Get("unconfTxs", opts)
	return res, err
}

// GetBlocksByTimestamp fetches an array of blocks for the time between the
// passed Epoch time millisecond timestamp and the beginning of the day
// Query pattern: `/blocks/:time?format=json(&api_code=:apiCode)`
func GetBlocksByTimestamp(timestamp int) (string, error) {
	beOpts := blockExplorerOpts{
		Opts:   api.Opts{},
		Format: "json",
	}
	opts := util.IntToString(timestamp) + "?" + client.EncodeOpts(beOpts)
	res, err := client.Get("blocks", opts)
	return res, err
}
