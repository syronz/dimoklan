package repo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dimoklan/consts/table"
	"dimoklan/model"
)

func (r *Repo) CreateUser(userRepo model.UserRepo) error {
	/*
		user.ID = consts.ParUser + user.ID
		user.SK = user.ID
		user.EntityType = consts.UserEntity
		av, err := dynamodbattribute.MarshalMap(user)
		if err != nil {
			return fmt.Errorf("error in marshmap user; err: %w", err)
		}

		input := &dynamodb.PutItemInput{
			Item:                av,
			TableName:           aws.String(consts.TableData),
			ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
		}

		if _, err = r.core.DynamoDB().PutItem(input); err != nil {
			return fmt.Errorf("put_item_failed_for_user; err:%w", err)
		}
	*/

	return nil
}

func (r *Repo) DeleteUser(userID string) error {
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

func (r *Repo) GetUserByEmail(email string) (model.UserRepo, error) {

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
