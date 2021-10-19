package util

import (
	"fmt"
	"strconv"
)

func GetDiffRatio(a float64, b float64) (diff float64) {
	value := (a - b) / b * 100
	diff, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return diff
}
