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

	fmt.Println(">>>>>", input)

	if _, err = bd.core.DynamoDB().PutItem(input); err != nil {
		return err
	}

	return nil
}

func (bd *BasDynamoDB) GetAllColors() (map[int]string, error) {

	return nil, nil
}
