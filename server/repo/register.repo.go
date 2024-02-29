package repo

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

func (r *Repo) CreateRegister(register types.Register) error {
	register.ActivationCode = consts.ParRegister + register.ActivationCode
	register.SK = register.ActivationCode
	register.EntityType = consts.RegisterEntity

	av, err := dynamodbattribute.MarshalMap(register)
	if err != nil {
		return fmt.Errorf("error in marshmap register; %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(consts.TableData),
		ConditionExpression: aws.String("attribute_not_exists(email)"),
	}

	_, err = r.core.DynamoDB().PutItem(input)
	if err != nil {
		var awsErr awserr.Error
		if errors.As(err, &awsErr) && awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
			return fmt.Errorf("email already registered")
		}
	}

	return err
}

func (r *Repo) ConfirmRegister(activationCode string) (types.Register, error) {
	activationCode = consts.ParRegister + activationCode

	params := &dynamodb.GetItemInput{
		TableName: aws.String(consts.TableData),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {S: aws.String(activationCode)},
			"SK": {S: aws.String(activationCode)},
		},
	}

	// read the item
	resp, err := r.core.DynamoDB().GetItem(params)
	if err != nil {
		return types.Register{}, fmt.Errorf("error in getting register entity; err: %w", err)
	}

	// unmarshal the dynamodb attribute values into a custom struct
	register := types.Register{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &register)
	if err != nil {
		return types.Register{}, fmt.Errorf("binding registration data failed; err: %w", err)
	}

	return register, nil
}
