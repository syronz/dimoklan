package types

import (
	"fmt"
	"strconv"
	"strings"

	"dimoklan/util"
)

type CELL string

func (c *CELL) GetX() int {
	nums := strings.Split(string(*c), ":")
	if len(nums) < 2 {
		return 0
	}

	num, _ := strconv.Atoi(nums[0])
	return num
}

func (c *CELL) GetY() int {
	nums := strings.Split(string(*c), ":")
	if len(nums) < 2 {
		return 0
	}

	num, _ := strconv.Atoi(nums[1])
	return num
}

func (c *CELL) ToString() string {
	return string(*c)
}

func (c *CELL) Set(x, y int) {
	*c = CELL(fmt.Sprintf("%v:%v", x, y))
}

func (c *CELL) ToFraction() string {
	x := util.CeilInt(float64(c.GetX()) / 10)
	y := util.CeilInt(float64(c.GetX()) / 10)

	return fmt.Sprintf("%d:%d", x, y)
}
