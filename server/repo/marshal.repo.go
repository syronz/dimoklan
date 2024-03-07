package repo

import (
	"context"

	"dimoklan/model"
)

func (r *Repo) CreateMarshal(ctx context.Context, marshalRepo model.MarshalRepo) error {
	/*
		marshal.UserID = consts.ParUser + marshal.UserID
		marshal.ID = consts.ParMarshal + marshal.ID

		fmt.Printf(">>>>>> M %+v\n", marshal)
		av, err := dynamodbattribute.MarshalMap(marshal)
		if err != nil {
			return fmt.Errorf("error in marshmap marshal; %w", err)
		}

		writeRequests := []dynamodb.WriteRequest{
			{PutRequest: &dynamodb.PutRequest{Item: av}},
		}

		// input := &dynamodb.PutItemInput{
		// 	Item:                av,
		// 	TableName:           aws.String(consts.TableData),
		// 	ConditionExpression: aws.String("attribute_not_exists(SK)"),
		// }

		input := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				consts.TableData: writeRequests,
			},
		}

		// _, err = r.core.DynamoDB().PutItem(input)
		_, err = r.core.DynamoDB().BatchWriteItem(input)
		if err != nil {
			// var awsErr awserr.Error
			// if errors.As(err, &awsErr) && awsErr.Code() == dynamodb.ErrCodeConditionalCheckFailedException {
			// 	return fmt.Errorf("marshal with this id already exist")
			// }
			return err
		}

		return err
	*/
	return nil
}

func (r *Repo) DeleteMarshal(ctx context.Context, userID, marshalID string) error {
	/*
		input := &dynamodb.DeleteItemInput{
			TableName: aws.String(consts.TableData),
			Key: map[string]*dynamodb.AttributeValue{
				"PK": {S: aws.String(userID)},
				"SK": {S: aws.String(marshalID)},
			},
		}

		if _, err := r.core.DynamoDB().DeleteItem(input); err != nil {
			return fmt.Errorf("delete_item_failed_for_marshal; err:%w", err)
		}
	*/

	return nil
}
