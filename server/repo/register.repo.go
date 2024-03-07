package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dimoklan/consts/table"
	"dimoklan/model"
)

func (r *Repo) CreateRegister(ctx context.Context, registerRepo model.RegisterRepo) error {
	item, err := attributevalue.MarshalMap(registerRepo)
	if err != nil {
		return fmt.Errorf("error in marshmap registerRepo; %w", err)
	}

	itemInput := &dynamodb.PutItemInput{
		TableName:           table.Data(),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	}

	_, err = r.core.DynamoDB().PutItem(ctx, itemInput)
	if err != nil {
		var conditionalCheckFailedErr *types.ConditionalCheckFailedException
		if errors.As(err, &conditionalCheckFailedErr) {
			return fmt.Errorf("email already exists")
		}

		return fmt.Errorf("error in register; err: %w", err)
	}
	return nil
}

func (r *Repo) ConfirmRegister(ctx context.Context, activationCode string) (model.RegisterRepo, error) {

	activationCodeMarshaled, err := attributevalue.Marshal(activationCode)
	if err != nil {
		return model.RegisterRepo{}, fmt.Errorf("error in marshal activation_code; err: %w", err)
	}

	params := &dynamodb.GetItemInput{
		TableName: table.Data(),
		Key: map[string]types.AttributeValue{
			"PK": activationCodeMarshaled,
			"SK": activationCodeMarshaled,
		},
	}

	// read the item
	resp, err := r.core.DynamoDB().GetItem(context.TODO(), params)
	if err != nil {
		return model.RegisterRepo{}, fmt.Errorf("error in getting register entity; err: %w", err)
	}

	// unmarshal the dynamodb attribute values into a custom struct
	registerRepo := model.RegisterRepo{}
	err = attributevalue.UnmarshalMap(resp.Item, &registerRepo)
	if err != nil {
		return model.RegisterRepo{}, fmt.Errorf("binding registration data failed; err: %w", err)
	}

	return registerRepo, nil

}
