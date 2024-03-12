package repo

import (
	"context"
	"fmt"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/model"
)

func (r *Repo) CreateAuth(ctx context.Context, auth model.Auth) error {

	return r.putUniqueItem(ctx, entity.Auth, auth.ToRepo())
	// item, err := attributevalue.MarshalMap(authRepo)
	// if err != nil {
	// 	return fmt.Errorf("error in marshmap authRepo; %w", err)
	// }

	// itemInput := &dynamodb.PutItemInput{
	// 	TableName:           table.Data(),
	// 	Item:                item,
	// 	ConditionExpression: aws.String("attribute_not_exists(PK)"),
	// }

	// _, err = r.core.DynamoDB().PutItem(ctx, itemInput)
	// if err != nil {
	// 	var conditionalCheckFailedErr *types.ConditionalCheckFailedException
	// 	if errors.As(err, &conditionalCheckFailedErr) {
	// 		return fmt.Errorf("email already exists")
	// 	}

	// 	return fmt.Errorf("error in auth; err: %w", err)
	// }
	// return nil
}

func (r *Repo) DeleteAuth(ctx context.Context, authID string) error {
	pk := hashtag.Auth + authID
	return r.deleteItem(ctx, entity.Auth, pk)
	// authIDMarshaled, err := attributevalue.Marshal(authID)
	// if err != nil {
	// 	return fmt.Errorf("error in marshal authID; err: %w", err)
	// }

	// params := &dynamodb.DeleteItemInput{
	// 	TableName: table.Data(),
	// 	Key: map[string]types.AttributeValue{
	// 		"PK": authIDMarshaled,
	// 		"SK": authIDMarshaled,
	// 	},
	// }

	// if _, err := r.core.DynamoDB().DeleteItem(ctx, params); err != nil {
	// 	return fmt.Errorf("delete item failed for auth; err:%w", err)
	// }

	// return nil
}

func (r *Repo) GetAuthByEmail(ctx context.Context, email string) (model.Auth, error) {
	authRepo := model.AuthRepo{}
	if err := r.getItem(ctx, entity.Auth, &authRepo, hashtag.Auth+email); err != nil {
		return model.Auth{}, err
	}
	fmt.Printf(">>>>>>> 1: %+v\n", authRepo.ToAPI())

	return authRepo.ToAPI(), nil

	// emailMarshaled, err := attributevalue.Marshal(hashtag.Auth + email)
	// if err != nil {
	// 	return model.AuthRepo{}, fmt.Errorf("error in marshal email; err: %w", err)
	// }

	// params := &dynamodb.GetItemInput{
	// 	TableName: table.Data(),
	// 	Key: map[string]types.AttributeValue{
	// 		"PK": emailMarshaled,
	// 		"SK": emailMarshaled,
	// 	},
	// }
	// resp, err := r.core.DynamoDB().GetItem(ctx, params)
	// if err != nil {
	// 	return model.AuthRepo{}, fmt.Errorf("error in getting auth entity; err: %w", err)
	// }

	// authRepo := model.AuthRepo{}
	// err = attributevalue.UnmarshalMap(resp.Item, &authRepo)
	// if err != nil {
	// 	return model.AuthRepo{}, fmt.Errorf("binding authRepo data failed; err: %w", err)
	// }

	// return authRepo, nil
}
