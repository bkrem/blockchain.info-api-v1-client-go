package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	chartType    = "transactions-per-second"
	dayTimespan  = "8days"
	weekTimespan = "5weeks"
)

func TestGetChartData(t *testing.T) {
	assert := assert.New(t)
	// JSON response
	jsonData, err := GetChartData(chartType, weekTimespan, "json")
	assert.Nil(err)
	assert.NotEmpty(jsonData)
	// CSV response
	csvData, err := GetChartData(chartType, weekTimespan, "csv")
	assert.Nil(err)
	assert.NotEmpty(csvData)
}

func TestGetPoolData(t *testing.T) {
	assert := assert.New(t)

	// no timespan
	poolData, err := GetPoolData("")
	assert.Nil(err)
	assert.NotEmpty(poolData)

	// with optional timespan
	poolData, err = GetPoolData(dayTimespan)
	assert.Nil(err)
	assert.NotEmpty(poolData)

	// invalid timespan
	poolData, err = GetPoolData(weekTimespan)
	assert.NotNil(err)
	assert.Empty(poolData)
}
