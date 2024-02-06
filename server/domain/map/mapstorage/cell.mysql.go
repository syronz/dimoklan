package mapstorage

import (
	_ "embed"
	"fmt"

	"dimoklan/internal/config"
	"dimoklan/types"
)

type CellMysql struct {
	core config.Core
}

func New(core config.Core) *CellMysql {
	return &CellMysql{
		core: core,
	}
}

//go:embed queries/create_cell.sql
var queryCreateCell string

func (ms *CellMysql) CreateCell(cell types.Cell) error {
	_, err := ms.core.BasicMasterDB().Exec(
		queryCreateCell,
		cell.X,
		cell.Y,
		cell.UserID,
		cell.Building,
		cell.Score,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ms *CellMysql) GetCellByCoord(x, y int) (types.Cell, error) {
	return types.Cell{}, nil
}

//go:embed queries/get_cell_user_ids.sql
var queryGetCellUsers string

func (ms *CellMysql) GetMapUsers(start types.Point, stop types.Point) (map[types.Point]int, error) {
	mapUsers := make(map[types.Point]int)

	rows, err := ms.core.BasicSlaveDB().Query(queryGetCellUsers, start.X, start.Y, stop.X, stop.Y)
	if err != nil {
		return mapUsers, err
	}
	defer rows.Close()

	for rows.Next() {
		var point types.Point
		var userID int
		if err := rows.Scan(&point.X, &point.Y, &userID); err != nil {
			return mapUsers, fmt.Errorf("error in scanning row for get map_user_id; %w", err)
		}

		mapUsers[point] = userID
	}

	return mapUsers, nil
}
