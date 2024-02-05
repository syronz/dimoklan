package types

type Point struct {
	X int
	Y int
}

type MapUsers struct {
	Points map[Point]int `json:"points"`
}

type MapColors struct {
	Points map[int]string `json:"points"`
}

type MapColor struct {
	Points [][]string `json:"points"`
}
