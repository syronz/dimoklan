package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dimoklan/consts/table"
	"dimoklan/model"
)

func (r *Repo) CreateUser(ctx context.Context, userRepo model.UserRepo) error {
	item, err := attributevalue.MarshalMap(userRepo)
	if err != nil {
		return fmt.Errorf("error in marshmap userRepo; %w", err)
	}

	itemInput := &dynamodb.PutItemInput{
		TableName:           table.Data(),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
	}

	_, err = r.core.DynamoDB().PutItem(ctx, itemInput)
	if err != nil {
		var conditionalCheckFailedErr *types.ConditionalCheckFailedException
		if errors.As(err, &conditionalCheckFailedErr) {
			return fmt.Errorf("user already exists")
		}

		return fmt.Errorf("error in user; err: %w", err)
	}
	return nil
}

func (r *Repo) DeleteUser(ctx context.Context, userID string) error {
	/*
		input := &dynamodb.DeleteItemInput{
			TableName: aws.String(consts.TableData),
			Key: map[string]*dynamodb.AttributeValue{
				"PK": {S: aws.String(userID)},
				"SK": {S: aws.String(userID)},
			},
		}

		if _, err := r.core.DynamoDB().DeleteItem(input); err != nil {
			return fmt.Errorf("delete_item_failed_for_user; err:%w", err)
		}
	*/

	return nil
}

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (model.UserRepo, error) {

	emailMarshaled, err := attributevalue.Marshal(email)
	if err != nil {
		return model.UserRepo{}, fmt.Errorf("error in marshal email; err: %w", err)
	}

	params := &dynamodb.GetItemInput{
		TableName: table.Data(),
		Key: map[string]types.AttributeValue{
			"PK": emailMarshaled,
			"SK": emailMarshaled,
		},
	}
	resp, err := r.core.DynamoDB().GetItem(context.TODO(), params)
	if err != nil {
		return model.UserRepo{}, fmt.Errorf("error in getting user entity; err: %w", err)
	}

	userRepo := model.UserRepo{}
	err = attributevalue.UnmarshalMap(resp.Item, &userRepo)
	if err != nil {
		return model.UserRepo{}, fmt.Errorf("binding userRepo data failed; err: %w", err)
	}

	return userRepo, nil

}
