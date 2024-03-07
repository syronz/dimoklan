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

func (r *Repo) CreateAuth(ctx context.Context, authRepo model.AuthRepo) error {
	item, err := attributevalue.MarshalMap(authRepo)
	if err != nil {
		return fmt.Errorf("error in marshmap authRepo; %w", err)
	}

	itemInput := &dynamodb.PutItemInput{
		TableName:           table.Data(),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	}

	_, err = r.core.DynamoDB().PutItem(ctx, itemInput)
	if err != nil {
		var conditionalCheckFailedErr *types.ConditionalCheckFailedException
		if errors.As(err, &conditionalCheckFailedErr) {
			return fmt.Errorf("email already exists")
		}

		return fmt.Errorf("error in auth; err: %w", err)
	}
	return nil

	/*
		auth.Email = consts.ParAuth + auth.Email
		auth.SK = auth.Email
		auth.EntityType = consts.AuthEntity
		av, err := dynamodbattribute.MarshalMap(auth)
		if err != nil {
			return fmt.Errorf("error in marshmap auth; err: %w", err)
		}

		input := &dynamodb.PutItemInput{
			Item:                av,
			TableName:           aws.String(consts.TableData),
			ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
		}

		if _, err = r.core.DynamoDB().PutItem(input); err != nil {
			return fmt.Errorf("put_item_failed_for_auth; err:%w", err)
		}

		return nil
	*/
	return nil
}

func (r *Repo) DeleteAuth(ctx context.Context, authID string) error {
	/*
		input := &dynamodb.DeleteItemInput{
			TableName: aws.String(consts.TableData),
			Key: map[string]*dynamodb.AttributeValue{
				"PK": {S: aws.String(authID)},
				"SK": {S: aws.String(authID)},
			},
		}

		if _, err := r.core.DynamoDB().DeleteItem(input); err != nil {
			return fmt.Errorf("delete_item_failed_for_auth; err:%w", err)
		}

	*/
	return nil
}

func (r *Repo) GetAuthByEmail(ctx context.Context, email string) (model.Auth, error) {
	/*
		email = consts.ParAuth + email

		params := &dynamodb.GetItemInput{
			TableName: aws.String(consts.TableData),
			Key: map[string]*dynamodb.AttributeValue{
				"PK": {
					S: aws.String(email),
				},
				"SK": {
					S: aws.String(email),
				},
			},
		}

		resp, err := r.core.DynamoDB().GetItem(params)
		if err != nil {
			return model.Auth{}, fmt.Errorf("error in getting auth entity; err: %w", err)
		}

		auth := model.Auth{}
		err = dynamodbattribute.UnmarshalMap(resp.Item, &auth)
		if err != nil {
			return model.Auth{}, fmt.Errorf("binding auth data failed; err: %w", err)
		}

		return auth, nil
	*/
	return model.Auth{}, nil
}
