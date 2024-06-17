package util

import (
	"math"
	"strings"

	"dimoklan/consts/hashtag"
)

func CeilInt(num float64) int {
	num = math.Ceil(num)

	return int(num)
}

func ExtractUserIDFromMarshalID(marshalID string) string {
	sections := strings.Split(marshalID[2:], ":")
	return hashtag.User + sections[0]
}
