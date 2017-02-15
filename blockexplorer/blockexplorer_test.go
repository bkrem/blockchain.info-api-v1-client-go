package blockexplorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sampleBlockHash   = "00000000000000000102abc71e79bacac2e9f6f22e6c42729c297f5e8b0cb1bf"
	sampleBlockIndex  = 1454914
	sampleBlockHeight = 123
	sampleAddress     = "1159LMHxYyvAfUUKXMP6ofTFQKhroS3enP"
	sampleHash160     = "00c89ed2a1c1991341a450eea89234929d803e47"
)

var (
	sampleMultiAddress = []string{"1159LMHxYyvAfUUKXMP6ofTFQKhroS3enP", "18BMm994cy3ovRnXyEYKzkLXeNgJrjhW6y"}
)

func TestGetBlock(t *testing.T) {
	assert := assert.New(t)
	blockByHash, err := GetBlockByHash(sampleBlockHash)
	blockByIndex, err2 := GetBlockByIndex(sampleBlockIndex)

	assert.Nil(err)
	assert.Nil(err2)
	assert.NotEmpty(blockByHash)
	assert.NotEmpty(blockByIndex)
	assert.Equal(blockByHash, blockByIndex) // TODO assert individually against .json file
}

func TestGetBlockHeight(t *testing.T) {
	assert := assert.New(t)
	blockHeight, err := GetBlockHeight(sampleBlockHeight)
	assert.Nil(err)
	assert.NotEmpty(blockHeight)
}

func TestGetLatestBlock(t *testing.T) {
	res, err := GetLatestBlock()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func TestGetAddress(t *testing.T) {
	assert := assert.New(t)
	res, err := GetAddress(sampleAddress, 0, 0)
	res2, err2 := GetAddress(sampleHash160, 0, 0)
	assert.Nil(err)
	assert.Nil(err2)
	assert.NotEmpty(res)
	assert.NotEmpty(res2)
	assert.Equal(res, res2) // TODO assert individually against .json file
}

func TestGetMultiAddress(t *testing.T) {
	res, err := GetMultiAddress(sampleMultiAddress)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}
