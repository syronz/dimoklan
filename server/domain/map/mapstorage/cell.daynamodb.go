package mapstorage

import (
	_ "embed"

	"dimoklan/internal/config"
	"dimoklan/types"
)

type CellDaynamo struct {
	core config.Core
}

func NewDaynamoCell(core config.Core) *CellDaynamo {
	return &CellDaynamo{
		core: core,
	}
}


func (ms *CellDaynamo) CreateCell(cell types.Cell) error {

	return nil
}

func (ms *CellDaynamo) GetCellByCoord(x, y int) (types.Cell, error) {
	return types.Cell{}, nil
}


func (ms *CellDaynamo) GetMapUsers(start types.Point, stop types.Point) (map[types.Point]int, error) {
	mapUsers := make(map[types.Point]int)


	return mapUsers, nil
}
