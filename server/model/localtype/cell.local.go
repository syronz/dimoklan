package localtype

import (
	"fmt"
	"strconv"
	"strings"

	"dimoklan/consts/hashtag"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/util"
)

type CELL struct {
	X   int `json:"x"`
	Y   int `json:"y"`
	str string
}

func NewCell(x, y int) (cell CELL) {
	cell.Set(x, y)
	return cell
}

func ParseCell(cellStr string) (CELL, error) {
	arr := strings.Split(cellStr, ":")

	if len(arr) < 3 || arr[0] != "c" {
		return CELL{}, fmt.Errorf("invalid pattern for cell not parsed; value: %v", cellStr)
	}

	x, err := strconv.Atoi(arr[1])
	if err != nil {
		return CELL{}, fmt.Errorf("x in cell pattern is not valid; err: %w", err)
	}

	y, err := strconv.Atoi(arr[2])
	if err != nil {
		return CELL{}, fmt.Errorf("x in cell pattern is not valid; err: %w", err)
	}
	return NewCell(x, y), nil
}

func ToCell(cellStr string) CELL {
	cell, _ := ParseCell(cellStr)
	return cell
}

func (c *CELL) IsEmpty() bool {
	return (c.X + c.Y) == 0
}

func (c *CELL) GetX() int {
	return c.X
}

func (c *CELL) GetY() int {
	return c.Y
}

func (c CELL) ToString() string {
	if c.str == "" {
		return fmt.Sprintf("%v%v:%v", hashtag.Cell, c.X, c.Y)
	}
	return c.str
}

func (c *CELL) Set(x, y int) {
	c.X = x
	c.Y = y
	c.str = fmt.Sprintf("%v%v:%v", hashtag.Cell, x, y)
}

func (c *CELL) SetStr() {
	c.str = fmt.Sprintf("%v%v:%v", hashtag.Cell, c.X, c.Y)
}

func (c *CELL) ToFraction() string {
	x := util.CeilInt(float64(c.GetX()) / 10)
	y := util.CeilInt(float64(c.GetY()) / 10)

	return fmt.Sprintf("%v%d:%d", hashtag.Fraction, x, y)
}

func (c *CELL) Validate() error {
	// nums := strings.Split(c.str, ":")
	// if len(nums) != 3 {
	// 	return fmt.Errorf("cell is not valid; code: %w", errstatus.ErrNotAcceptable)
	// }
	if c.X+c.Y == 0 {
		return fmt.Errorf("cell is not valid; code: %w", errstatus.ErrNotAcceptable)
	}

	return nil
}

// MarshalBinary is used to save CELL types to Redis
func (c CELL) MarshalBinary() ([]byte, error) {
	return []byte(c.str), nil
}

/*
func (s *testScanSliceStruct) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, s)
}
*/

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (c *CELL) UnmarshalBinary(data []byte) error {
	fmt.Printf(">>>>>>> 1999: %+v\n", data)

	str := string(data)
	parts := strings.Split(str, ":")
	if len(parts) != 3 || parts[0] != "c" {
		return fmt.Errorf("invalid cell format: %s", str)
	}

	var row, col int
	if _, err := fmt.Sscanf(parts[1]+":"+parts[2], "%d:%d", &row, &col); err != nil {
		return err
	}

	c.X = row
	c.Y = col
	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// func (c *CELL) UnmarshalText(text []byte) error {
// 	fmt.Printf(">>>>>>> ole 999: %+v\n", string(text))

// 	parts := strings.Split(string(text), ":")
// 	if len(parts) != 2 {
// 		return fmt.Errorf("invalid cell format")
// 	}

// 	var row, col int
// 	_, err := fmt.Sscanf(parts[1], "%d:%d", &row, &col)
// 	if err != nil {
// 		return err
// 	}

// 	c.X = row
// 	c.Y = col
// 	return nil
// }
