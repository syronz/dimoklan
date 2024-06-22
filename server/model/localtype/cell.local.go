package localtype

import (
	"fmt"
	"strconv"
	"strings"

	"dimoklan/consts/hashtag"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/util"
)

type CELL string

func (c *CELL) GetX() int {
	nums := strings.Split(string(*c), ":")
	if len(nums) < 3 {
		return 0
	}

	num, _ := strconv.Atoi(nums[1])
	return num
}

func (c *CELL) GetY() int {
	nums := strings.Split(string(*c), ":")
	if len(nums) < 3 {
		return 0
	}

	num, _ := strconv.Atoi(nums[2])
	return num
}

func (c *CELL) ToString() string {
	return string(*c)
}

func (c *CELL) Set(x, y int) {
	*c = CELL(fmt.Sprintf("%v%v:%v", hashtag.Cell, x, y))
}

func (c *CELL) ToFraction() string {
	x := util.CeilInt(float64(c.GetX()) / 10)
	y := util.CeilInt(float64(c.GetY()) / 10)

	return fmt.Sprintf("%v%d:%d", hashtag.Fraction, x, y)
}

func (c *CELL) Validate() error {
	nums := strings.Split(string(*c), ":")
	if len(nums) != 3 {
		return fmt.Errorf("cell is not valid; code: %w", errstatus.ErrNotAcceptable)
	}

	return nil
}
