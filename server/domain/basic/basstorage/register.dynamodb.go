package basstorage

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/consts"
	"dimoklan/types"
)

func (bd *BasDynamoDB) CreateRegister(register types.Register) error {
	av, err := dynamodbattribute.MarshalMap(register)
	if err != nil {
		return fmt.Errorf("error in marshmap register; %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(consts.TableRegister),
		ConditionExpression: aws.String("attribute_not_exists(email)"),
	}

	_, err = bd.core.DynamoDB().PutItem(input)
	if err != nil {
		var awsErr awserr.Error
		if errors.As(err, &awsErr) && awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
			return fmt.Errorf("email already registered")
		}
	}

	return err
}

func (bd *BasDynamoDB) ConfirmRegister(activationCode string) (types.Register, error) {
	// Define the query input
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(consts.TableRegister), // Replace with your table name
		IndexName:              aws.String("activation_code_index"),
		KeyConditionExpression: aws.String("activation_code = :b"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":b": {
				S: aws.String(activationCode),
			},
		},
		// ProjectionExpression:   ,
	}

	// Perform the query
	result, err := bd.core.DynamoDB().Query(queryInput)
	if err != nil {
		return types.Register{}, fmt.Errorf("error in getting register entity; err: %w", err)
	}

	var registers []types.Register
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &registers)
	if err != nil {
		return types.Register{}, fmt.Errorf("binding registration data failed; err: %w", err)
	}

	if len(registers) == 0 {
		return types.Register{}, errors.New("activation code has been expired")
	}

	return registers[0], nil
}
