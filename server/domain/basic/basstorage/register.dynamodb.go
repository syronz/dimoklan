package basstorage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/types"
)

func (bd *BasDynamoDB) CreateRegister(register types.Register) error {
	av, err := dynamodbattribute.MarshalMap(register)
	if err != nil {
		return fmt.Errorf("error in marshmap register; %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: bd.registerTable,
	}

	_, err = bd.core.DynamoDB().PutItem(input)

	return err
}
