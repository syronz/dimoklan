package basstorage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/consts"
	"dimoklan/types"
)

func (bd *BasDynamoDB) GetUserByColor(color string) (types.User, error) {

	return types.User{}, nil
}

func (bd *BasDynamoDB) CreateUser(user types.User) error {
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("error in marshmap user; err: %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(consts.TableUser),
		ConditionExpression: aws.String("attribute_not_exists(email)"),
	}

	if _, err = bd.core.DynamoDB().PutItem(input); err != nil {
		return err
	}

	return nil
}

func (bd *BasDynamoDB) GetAllColors() (map[int]string, error) {

	return nil, nil
}

func (bd *BasDynamoDB) GetUserByEmail(email string) (types.User, error) {
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(consts.TableUser),
		IndexName:              aws.String(consts.IndexEmail),
		KeyConditionExpression: aws.String("email = :b"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":b": {
				S: aws.String(email),
			},
		},
	}

	result, err := bd.core.DynamoDB().Query(queryInput)
	if err != nil {
		return types.User{}, fmt.Errorf("error in getting user entity; err: %w", err)
	}

	var users []types.User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	fmt.Println(">>>>>>", users)
	if err != nil {
		return types.User{}, fmt.Errorf("binding user data failed; err: %w", err)
	}

	if len(users) > 0 {
		return users[0], nil
	}

	return types.User{}, nil
}
