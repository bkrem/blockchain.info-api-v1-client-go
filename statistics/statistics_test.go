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
	data, err := GetChartData(chartType, timespan)
	assert.Nil(t, err)
	assert.NotEmpty(t, data)
}
