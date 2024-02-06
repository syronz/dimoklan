package mapstorage

import "dimoklan/types"

type MapStorage interface {
	CreateCell(types.Cell) error
	GetCellByCoord(int, int) (types.Cell, error)
	GetMapUsers(types.Point, types.Point) (map[types.Point]int, error)
}

