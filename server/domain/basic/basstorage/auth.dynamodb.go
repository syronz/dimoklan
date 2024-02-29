package basstorage

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/consts"
	"dimoklan/types"
)

func (bd *BasDynamoDB) CreateAuth(auth types.Auth) error {
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

	if _, err = bd.core.DynamoDB().PutItem(input); err != nil {
		return fmt.Errorf("put_item_failed_for_auth; err:%w", err)
	}

	return nil
}

func (bd *BasDynamoDB) GetAuthByEmail(email string) (types.Auth, error) {
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
		return types.Auth{}, fmt.Errorf("error in getting auth entity; err: %w", err)
	}

	auth := types.Auth{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &auth)
	if err != nil {
		return types.Auth{}, fmt.Errorf("binding auth data failed; err: %w", err)
	}

	return auth, nil
}
