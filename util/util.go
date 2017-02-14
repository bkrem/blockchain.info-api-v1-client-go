package util

import (
	"regexp"
	"strconv"
)

func StringToFloat64(floatString string) (float64, error) {
	// Remove comma separators in response to parse as float
	formattedRes := regexp.MustCompile(",").ReplaceAllString(floatString, "")
	// Cast the formatted string to float64
	parsedRes, err := strconv.ParseFloat(formattedRes, 64)
	return parsedRes, err
}

func Float64ToString(floatVal float64) string {
	return strconv.FormatFloat(floatVal, 'f', -1, 64)
}

func IntToString(intVal int) string {
	return strconv.Itoa(intVal)
}
