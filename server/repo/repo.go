package repo

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dimoklan/consts"
	"dimoklan/consts/table"
	"dimoklan/internal/config"
	"dimoklan/model/localtype"
)

type Repo struct {
	core config.Core
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

// putUniqueItems is common function for adding items
// func (r *Repo) putItems(ctx context.Context, entityType string, itemRepos any) error {
// 	// Check if items is a slice
// 	sliceValue := reflect.ValueOf(itemRepos)
// 	if sliceValue.Kind() != reflect.Slice {
// 		return fmt.Errorf("items must be a slice")
// 	}

// 	// Iterate over the items in the slice
// 	for i := 0; i < sliceValue.Len(); i++ {
// 		itemRepo := sliceValue.Index(i).Interface()

// 		item, err := attributevalue.MarshalMap(itemRepo)
// 		if err != nil {
// 			return fmt.Errorf("error in marshalmap item; entity: %v; %v", entityType, err)
// 		}

// 		itemInput := &dynamodb.PutItemInput{
// 			TableName: table.Data(),
// 			Item:      item,
// 		}

// 		_, err = r.core.DynamoDB().PutItem(ctx, itemInput)
// 		if err != nil {
// 			return fmt.Errorf("error in putting item; entity: %v; err: %w", entityType, err)
// 		}
// 	}

// 	return nil
// }

func (r *Repo) putItems(ctx context.Context, entityType string, itemRepos interface{}) error {
	// Check if items is a slice
	sliceValue := reflect.ValueOf(itemRepos)
	if sliceValue.Kind() != reflect.Slice {
		return fmt.Errorf("items must be a slice")
	}

	// Prepare the list of WriteRequests for BatchWriteItem
	writeRequests := make([]types.WriteRequest, 0, sliceValue.Len())

	// Iterate over the items in the slice
	for i := 0; i < sliceValue.Len(); i++ {
		itemRepo := sliceValue.Index(i).Interface()

		item, err := attributevalue.MarshalMap(itemRepo)
		if err != nil {
			return fmt.Errorf("error in marshalmap item; entity: %v; %v", entityType, err)
		}

		writeRequests = append(writeRequests, types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
			},
		})
	}

	// Perform BatchWriteItem operation
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			consts.TableData: writeRequests,
		},
	}

	_, err := r.core.DynamoDB().BatchWriteItem(ctx, input)
	if err != nil {
		return fmt.Errorf("error in batch writing items; entity: %v; err: %w", entityType, err)
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

func (r *Repo) deleteItems(ctx context.Context, items []localtype.Delete) error {
	writeRequests := make([]types.WriteRequest, len(items))
	for i := range items {

		pkMarshaled, err := attributevalue.Marshal(items[i].PK)
		if err != nil {
			return fmt.Errorf("error in marshal pk for delete; pk: %v ; err: %w", items[i].PK, err)
		}

		skMarshaled, err := attributevalue.Marshal(items[i].SK)
		if err != nil {
			return fmt.Errorf("error in marshal sk for delete; sk: %v ; err: %w", items[i].SK, err)
		}

		writeRequests[i] = types.WriteRequest{
			DeleteRequest: &types.DeleteRequest{
				Key: map[string]types.AttributeValue{
					"PK": pkMarshaled,
					"SK": skMarshaled,
				},
			},
		}
	}

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			consts.TableData: writeRequests,
		},
	}

	_, err := r.core.DynamoDB().BatchWriteItem(ctx, input)
	if err != nil {
		return fmt.Errorf("error in batch delete items; err: %w", err)
	}

	return nil
}

func (r *Repo) getItem(ctx context.Context, entityType string, item any, pk string, sk ...string) error {
	pkMarshaled, err := attributevalue.Marshal(pk)
	if err != nil {
		return fmt.Errorf("error in marshal pk; entity: %v; err: %w", entityType, err)
	}

	skMarshaled := pkMarshaled
	if len(sk) > 0 {
		skMarshaled, err = attributevalue.Marshal(sk[0])
		if err != nil {
			return fmt.Errorf("error in marshal sk; entity: %v; err: %w", entityType, err)
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
		return fmt.Errorf("error in getting item; entity: %v; err: %w", entityType, err)
	}

	err = attributevalue.UnmarshalMap(resp.Item, item)
	if err != nil {
		return fmt.Errorf("binding item data failed; err: %w", err)
	}

	return nil
}
