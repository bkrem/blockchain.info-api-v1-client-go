package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	chartType = "transactions-per-second"
	timespan  = "5weeks"
)

func TestGetChartData(t *testing.T) {
	assert := assert.New(t)
	// JSON response
	jsonData, err := GetChartData(chartType, timespan, "json")
	assert.Nil(err)
	assert.NotEmpty(jsonData)
	// CSV response
	csvData, err := GetChartData(chartType, timespan, "csv")
	assert.Nil(err)
	assert.NotEmpty(csvData)
}
