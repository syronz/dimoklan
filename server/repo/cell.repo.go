package repo

import (
	_ "embed"

	"dimoklan/model"
)

func (r *Repo) CreateCell(cell model.Cell) error {
	/*
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

		if _, err = r.core.DynamoDB().PutItem(input); err != nil {
			return fmt.Errorf("put_item_failed_for_cell; err:%w", err)
		}

		return err
	*/
	return nil
}

func (r *Repo) GetCellByCoord(x, y int) (model.Cell, error) {
	return model.Cell{}, nil
}

func (r *Repo) GetMapUsers(start model.Point, stop model.Point) (map[model.Point]int, error) {
	mapUsers := make(map[model.Point]int)

	return mapUsers, nil
}
