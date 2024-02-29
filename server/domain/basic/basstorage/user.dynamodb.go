package basstorage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/consts"
	"dimoklan/types"
)

// func (bd *BasDynamoDB) CreateUser(user types.User) error {
// 	av, err := dynamodbattribute.MarshalMap(user)
// 	if err != nil {
// 		return fmt.Errorf("error in marshmap user; err: %w", err)
// 	}

// 	input := &dynamodb.PutItemInput{
// 		Item:                av,
// 		TableName:           aws.String(consts.TableUser),
// 		ConditionExpression: aws.String("attribute_not_exists(email)"),
// 	}

// 	if _, err = bd.core.DynamoDB().PutItem(input); err != nil {
// 		return err
// 	}

// 	return nil
// }

func (bd *BasDynamoDB) CreateUser(user types.User) error {
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

	if _, err = bd.core.DynamoDB().PutItem(input); err != nil {
		return fmt.Errorf("put_item_failed_for_user; err:%w", err)
	}

	return nil
}

func (bd *BasDynamoDB) DeleteUser(userID string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(consts.TableData),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {S: aws.String(userID)},
			"SK": {S: aws.String(userID)},
		},
	}

	if _, err := bd.core.DynamoDB().DeleteItem(input); err != nil {
		return fmt.Errorf("delete_item_failed_for_user; err:%w", err)
	}

	return nil
}

// func (bd *BasDynamoDB) GetUserByEmail(email string) (types.User, error) {
// 	queryInput := &dynamodb.QueryInput{
// 		TableName:              aws.String(consts.TableUser),
// 		IndexName:              aws.String(consts.IndexEmail),
// 		KeyConditionExpression: aws.String("email = :b"),
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":b": {
// 				S: aws.String(email),
// 			},
// 		},
// 	}

// 	result, err := bd.core.DynamoDB().Query(queryInput)
// 	if err != nil {
// 		return types.User{}, fmt.Errorf("error in getting user entity; err: %w", err)
// 	}

// 	var users []types.User
// 	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
// 	fmt.Println(">>>>>>", users)
// 	if err != nil {
// 		return types.User{}, fmt.Errorf("binding user data failed; err: %w", err)
// 	}

// 	if len(users) > 0 {
// 		return users[0], nil
// 	}

// 	return types.User{}, nil
// }

func (bd *BasDynamoDB) GetUserByEmail(email string) (types.User, error) {
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

	resp, err := bd.core.DynamoDB().GetItem(params)
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
