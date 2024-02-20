package mapstorage

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"dimoklan/internal/config"
	"dimoklan/types"
)

const cellTable = "cells"

type CellDaynamo struct {
	core      config.Core
	cellTable *string
}

func NewDaynamoCell(core config.Core) *CellDaynamo {
	return &CellDaynamo{
		core:      core,
		cellTable: aws.String(cellTable),
	}
}

func (ms *CellDaynamo) CreateCell(cell types.Cell) error {
	item := map[string]*dynamodb.AttributeValue{
		"cell": {
			S: aws.String(fmt.Sprintf("%v:%v", cell.X, cell.Y)),
		},
		"user_id": {
			N: aws.String(strconv.Itoa(cell.UserID)),
		},
		"score": {
			N: aws.String(strconv.Itoa(cell.UserID)),
		},
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: ms.cellTable,
	}

	_, err := ms.core.DynamoDB().PutItem(input)

	return err
}

func (ms *CellDaynamo) GetCellByCoord(x, y int) (types.Cell, error) {
	return types.Cell{}, nil
}

func (ms *CellDaynamo) GetMapUsers(start types.Point, stop types.Point) (map[types.Point]int, error) {
	mapUsers := make(map[types.Point]int)

	return mapUsers, nil
}
