package repo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/model"
	"dimoklan/model/localtype"
)

func (r *Repo) CreateMarshal(ctx context.Context, marshal model.Marshal) error {
	marshalInFraction := model.MarshalRepo{
		PK:         hashtag.Fraction + marshal.Cell.ToFraction(),
		SK:         hashtag.MarshalEx + marshal.ID,
		EntityType: entity.MarshalExist,
		Cell:       marshal.Cell,
		CreatedAt:  time.Now().Unix(),
	}

	marshals := []model.MarshalRepo{
		marshal.ToRepo(),
		marshalInFraction,
	}

	return r.putItems(ctx, entity.Marshal, marshals)
	// item, err := attributevalue.MarshalMap(marshal)
	// if err != nil {
	// 	return fmt.Errorf("error in marshmap marshal; %w", err)
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
	// 		return fmt.Errorf("marshal already exists")
	// 	}

	// 	return fmt.Errorf("error in marshal; err: %w", err)
	// }
	// return nil
}

func (r *Repo) DeleteMarshal(ctx context.Context, userID, marshalID, fraction string) error {
	batchDelete := make([]localtype.Delete, 2)
	batchDelete[0].PK = hashtag.User + userID
	batchDelete[0].SK = hashtag.Marshal + marshalID

	batchDelete[1].PK = hashtag.Fraction + fraction
	batchDelete[1].SK = hashtag.MarshalEx + marshalID
	return r.deleteItems(ctx, batchDelete)

	// userIDMarshaled, err := attributevalue.Marshal(userID)
	// if err != nil {
	// 	return fmt.Errorf("error in marshal userID; err: %w", err)
	// }

	// marshalIDMarshaled, err := attributevalue.Marshal(marshalID)
	// if err != nil {
	// 	return fmt.Errorf("error in marshal marshalID; err: %w", err)
	// }

	// params := &dynamodb.DeleteItemInput{
	// 	TableName: table.Data(),
	// 	Key: map[string]types.AttributeValue{
	// 		"PK": userIDMarshaled,
	// 		"SK": marshalIDMarshaled,
	// 	},
	// }

	// if _, err := r.core.DynamoDB().DeleteItem(ctx, params); err != nil {
	// 	return fmt.Errorf("delete item failed for marshal; err:%w", err)
	// }

	return nil
}

func (r *Repo) GetMarshal(ctx context.Context, id string) (model.Marshal, error) {

	parsedID := strings.Split(id, ":")
	if len(parsedID) != 2 {
		return model.Marshal{}, fmt.Errorf("id is not valid; code:%w", errstatus.ErrNotAcceptable)
	}

	marshalRepo := model.MarshalRepo{}
	err := r.getItem(ctx,
		entity.Marshal,
		&marshalRepo,
		hashtag.User+parsedID[0],
		hashtag.Marshal+id)

	if err != nil {
		return model.Marshal{}, err
	}

	return marshalRepo.ToAPI(), nil
}