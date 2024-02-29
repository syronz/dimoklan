package mapstorage

import (
	_ "embed"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/consts"
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
	cell.Fraction = consts.ParFraction + cell.Fraction
	cell.Cell = consts.ParCell + cell.Cell
	cell.EntityType = consts.CellEntity

	cellAV, err := dynamodbattribute.MarshalMap(cell)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(consts.TableData),
		Item:                cellAV,
		ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
	}

	if _, err = ms.core.DynamoDB().PutItem(input); err != nil {
		return fmt.Errorf("put_item_failed_for_cell; err:%w", err)
	}

	return err
}

func (ms *CellDaynamo) GetCellByCoord(x, y int) (types.Cell, error) {
	return types.Cell{}, nil
}

func (ms *CellDaynamo) GetMapUsers(start types.Point, stop types.Point) (map[types.Point]int, error) {
	mapUsers := make(map[types.Point]int)

	return mapUsers, nil
}
