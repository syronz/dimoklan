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
	"dimoklan/internal/config"
)

type Repo struct {
	core          config.Core
	registerTable *string
}

func NewRepo(core config.Core) *Repo {
	return &Repo{
		core: core,
	}
}

// putUniqueItem is common function for adding an item
func (r *Repo) putUniqueItem(ctx context.Context, entityType string, itemRepo any) error {
	item, err := attributevalue.MarshalMap(itemRepo)
	if err != nil {
		return fmt.Errorf("error in marshalmap item; entity: %v; %v", entityType, err)
	}

	itemInput := &dynamodb.PutItemInput{
		TableName:           table.Data(),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
	}

	_, err = r.core.DynamoDB().PutItem(ctx, itemInput)
	if err != nil {
		var conditionalCheckFailedErr *types.ConditionalCheckFailedException
		if errors.As(err, &conditionalCheckFailedErr) {
			return fmt.Errorf("item already exists; entity: %v; item: %v", entityType, itemRepo)
		}

		return fmt.Errorf("error in putting item; entity: %v; err: %w", entityType, err)
	}

	return nil
}


// deleteItem is a common function for deleting an item
func (r *Repo) deleteItem(ctx context.Context, entityType, pk string, sk ...string) error {
	pkMarshaled, err := attributevalue.Marshal(pk)
	if err != nil {
		return fmt.Errorf("error in marshal pk; entity: %v ; err: %w", entityType, err)
	}

	skMarshaled := pkMarshaled
	if len(sk) > 0 {
		skMarshaled, err = attributevalue.Marshal(sk)
		if err != nil {
			return fmt.Errorf("error in marshal sk; entity: %v ; err: %w", entityType, err)
		}
	}

	params := &dynamodb.DeleteItemInput{
		TableName: table.Data(),
		Key: map[string]types.AttributeValue{
			"PK": pkMarshaled,
			"SK": skMarshaled,
		},
	}

	if _, err := r.core.DynamoDB().DeleteItem(ctx, params); err != nil {
		return fmt.Errorf("delete item failed; entity: %v; err:%w", entityType, err)
	}

	return nil
}


func (r *Repo) getItem(ctx context.Context, entityType string, item any, pk string, sk ...string) error {
	pkMarshaled, err := attributevalue.Marshal(pk)
	if err != nil {
		return fmt.Errorf("error in marshal pk; entity: %v ; err: %w", entityType, err)
	}

	skMarshaled := pkMarshaled
	if len(sk) > 0 {
		skMarshaled, err = attributevalue.Marshal(sk)
		if err != nil {
			return fmt.Errorf("error in marshal sk; entity: %v ; err: %w", entityType, err)
		}
	}
	

	params := &dynamodb.GetItemInput{
		TableName: table.Data(),
		Key: map[string]types.AttributeValue{
			"PK": pkMarshaled,
			"SK": skMarshaled,
		},
	}
	resp, err := r.core.DynamoDB().GetItem(ctx, params)
	if err != nil {
		return fmt.Errorf("error in getting auth entity; err: %w", err)
	}

	err = attributevalue.UnmarshalMap(resp.Item, item)
	if err != nil {
		return fmt.Errorf("binding item data failed; err: %w", err)
	}

	return nil
}
