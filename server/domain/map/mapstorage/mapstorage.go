package mapstorage

import "dimoklan/types"

type MapStorage interface {
	CreateCell(types.Cell) error
	GetCellByCoord(int, int) (types.Cell, error)
	GetMap(types.Cell, types.Cell) ([]types.Cell, error)
}

