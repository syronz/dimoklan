package mapstorage

import (
	_ "embed"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/internal/config"
	"dimoklan/types"
)

const mapTable = "map"

type CellDaynamo struct {
	core      config.Core
	mapTable *string
}

func NewDaynamoCell(core config.Core) *CellDaynamo {
	return &CellDaynamo{
		core:      core,
		mapTable: aws.String(mapTable),
	}
}

func (ms *CellDaynamo) CreateCell(cell types.Cell) error {
	type Item struct {
		Fraction   string `json:"fraction"`
		Cell       string `json:"cell"`
		UserID     int    `json:"user_id"`
		Score      int    `json:"score"`
		LastUpdate int64  `json:"last_update"`
	}

	item := Item{
		Fraction:   cell.Fraction,
		Cell:       cell.Cell,
		UserID:     cell.UserID,
		Score:      cell.Score,
		LastUpdate: cell.LastUpdate,
	}

	// Marshal the item into a Map of attribute values
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: ms.mapTable,
	}

	_, err = ms.core.DynamoDB().PutItem(input)

	return err
}

func (ms *CellDaynamo) GetCellByCoord(x, y int) (types.Cell, error) {
	return types.Cell{}, nil
}

func (ms *CellDaynamo) GetMapUsers(start types.Point, stop types.Point) (map[types.Point]int, error) {
	mapUsers := make(map[types.Point]int)

	return mapUsers, nil
}
