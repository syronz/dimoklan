package repo

import (
	"context"
	_ "embed"

	"dimoklan/consts/entity"
	"dimoklan/model"
)

func (r *Repo) CreateCell(ctx context.Context, cell model.Cell) error {
	return r.putUniqueItem(ctx, entity.Cell, cell.ToRepo())
	// item, err := attributevalue.MarshalMap(cellRepo)
	// if err != nil {
	// 	return fmt.Errorf("error in marshmap cellRepo; %w", err)
	// }

	// itemInput := &dynamodb.PutItemInput{
	// 	TableName:           table.Data(),
	// 	Item:                item,
	// 	ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
	// }

	// _, err = r.core.DynamoDB().PutItem(ctx, itemInput)
	// if err != nil {
	// 	var conditionalCheckFailedErr *types.ConditionalCheckFailedException
	// 	if errors.As(err, &conditionalCheckFailedErr) {
	// 		return fmt.Errorf("cell already exists")
	// 	}

	// 	return fmt.Errorf("error in cell; err: %w", err)
	// }
	// return nil
}

func (r *Repo) GetCellByCoord(ctx context.Context, x, y int) (model.Cell, error) {
	return model.Cell{}, nil
}

func (r *Repo) GetMapUsers(ctx context.Context, start model.Point, stop model.Point) (map[model.Point]int, error) {
	mapUsers := make(map[model.Point]int)

	return mapUsers, nil
}
