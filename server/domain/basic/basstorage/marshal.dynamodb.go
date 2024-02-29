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

func (bd *BasDynamoDB) CreateMarshal(marshal types.Marshal) error {
	marshal.UserID = consts.ParUser + marshal.UserID
	marshal.ID = consts.ParMarshal + marshal.ID


	fmt.Printf(">>>>>> M %+v\n", marshal)
	av, err := dynamodbattribute.MarshalMap(marshal)
	if err != nil {
		return fmt.Errorf("error in marshmap marshal; %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(consts.TableData),
		ConditionExpression: aws.String("attribute_not_exists(SK)"),
	}

	_, err = bd.core.DynamoDB().PutItem(input)
	if err != nil {
		var awsErr awserr.Error
		if errors.As(err, &awsErr) && awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
			return fmt.Errorf("marshal with this id already exist")
		}
	}

	return err
}

func (bd *BasDynamoDB) DeleteMarshal(userID, marshalID string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(consts.TableData),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {S: aws.String(userID)},
			"SK": {S: aws.String(marshalID)},
		},
	}

	if _, err := bd.core.DynamoDB().DeleteItem(input); err != nil {
		return fmt.Errorf("delete_item_failed_for_marshal; err:%w", err)
	}

	return nil
}

