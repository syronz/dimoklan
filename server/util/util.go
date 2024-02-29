package util

import "math"

func CeilInt(num float64) int {
	num = math.Ceil(num)

	return int(num)
}
