package repo

import (
	"context"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/model"
)

func (r *Repo) CreateMarshal(ctx context.Context, marshal model.Marshal) error {
	return r.putUniqueItem(ctx, entity.Marshal, marshal.ToRepo())
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

func (r *Repo) DeleteMarshal(ctx context.Context, userID, marshalID string) error {
	pk := hashtag.User + userID
	sk := hashtag.Marshal + marshalID
	return r.deleteItem(ctx, entity.Marshal, pk, sk)

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
