package repo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/consts"
	"dimoklan/types"
)

func (r *Repo) CreateUser(user types.User) error {
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

	return nil
}

func (r *Repo) DeleteUser(userID string) error {
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

	return nil
}

func (r *Repo) GetUserByEmail(email string) (types.User, error) {
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
		return types.User{}, fmt.Errorf("error in getting user entity; err: %w", err)
	}

	user := types.User{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &user)
	if err != nil {
		return types.User{}, fmt.Errorf("binding user data failed; err: %w", err)
	}

	return user, nil
}
