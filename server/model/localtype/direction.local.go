package localtype

import (
	"fmt"
	"slices"

	"dimoklan/consts/gp"
)

type DIRECTION string

func SetDirection(direction string) (DIRECTION, error) {
	if !slices.Contains(gp.Directions(), direction) {
		return "", fmt.Errorf("direction is not valid; direction: %v", direction)
	}

	return DIRECTION(direction), nil
}
