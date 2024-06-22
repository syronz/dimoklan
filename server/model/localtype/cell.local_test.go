package localtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFraction(t *testing.T) {
	cell_3_9 := CELL("")
	cell_3_9.Set(3, 9)
	samples := map[string]struct {
		cell CELL
		out  string
	}{
		"cell_3_9_to_fraction_from_CELL": {
			cell: CELL("c:3:9"),
			out:  "f:1:1",
		},
		"cell_3_9_to_fraction_from_set": {
			cell: cell_3_9,
			out:  "f:1:1",
		},
		"cell_1_1_to_fraction_from_CELL": {
			cell: CELL("c:1:1"),
			out:  "f:1:1",
		},
		"cell_12_6_to_fraction_from_CELL": {
			cell: CELL("c:12:6"),
			out:  "f:2:1",
		},
	}

	for testName, sample := range samples {
		t.Run(testName, func(t *testing.T) {
			result := sample.cell.ToFraction()
			assert.Equal(t, sample.out, result)

			result = sample.cell.ToFraction()
			assert.Equal(t, sample.out, result)
		})
	}
}
