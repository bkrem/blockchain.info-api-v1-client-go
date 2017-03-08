package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const floatType float64 = 0.1

func TestGetTicker(t *testing.T) {
	ticker, err := GetTicker()
	assert.Nil(t, err)
	assert.NotEmpty(t, ticker)
}

func TestFromBTC(t *testing.T) {
	assert := assert.New(t)
	val, err := FromBTC(100000, "USD")
	assert.Nil(err)
	assert.NotEmpty(val)
	assert.IsType(floatType, val)
}

func TestToBTC(t *testing.T) {
	assert := assert.New(t)
	val, err := ToBTC(100000.123, "USD")
	assert.Nil(err)
	assert.NotEmpty(val)
	assert.IsType(floatType, val)
}
