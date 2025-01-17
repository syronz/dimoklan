package repo

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dimoklan/consts/hashtag"
	"dimoklan/consts/table"
	"dimoklan/model"
)

func (r *Repo) GetFractions(ctx context.Context, coordinates []string) ([]model.Fraction, error) {
	fractions := make([]model.Fraction, 0, 100)

	for _, coord := range coordinates {
		coordinate := hashtag.Fraction + coord

		var err error
		var response *dynamodb.QueryOutput
		keyEx := expression.Key("PK").Equal(expression.Value(coordinate))
		expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
		if err != nil {
			return nil, fmt.Errorf("building expression failed for get-fraction query. err: %w", err)
		}

		queryPaginator := dynamodb.NewQueryPaginator(r.core.DynamoDB(), &dynamodb.QueryInput{
			TableName:                 table.Data(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		})
		for queryPaginator.HasMorePages() {
			response, err = queryPaginator.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to query for get-fractions. err: %w", err)
			}
			var fractionPage []model.Fraction
			err = attributevalue.UnmarshalListOfMaps(response.Items, &fractionPage)
			if err != nil {
				return nil, fmt.Errorf("couldn't unmarshal response for fractions. err: %w", err)
			}
			fractions = append(fractions, fractionPage...)
		}

	}

	return fractions, nil
}

func (r *Repo) UpdateFraction(ctx context.Context, fraction model.Fraction) error {
	var err error
	var response *dynamodb.UpdateItemOutput
	var attributeMap map[string]map[string]interface{}
	update := expression.Set(expression.Name("EntityType"), expression.Value(fraction.EntityType))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return fmt.Errorf("couldn't build expression for update fraction. err: %w", err)
	}

	response, err = r.core.DynamoDB().UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 table.Data(),
		Key:                       fraction.GetKey(r.core),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})
	if err != nil {
		return fmt.Errorf("couldn't update fraction; cell:%v; err: %w", fraction.CellStr, err)
	}

	err = attributevalue.UnmarshalMap(response.Attributes, &attributeMap)
	if err != nil {
		return fmt.Errorf("couldn't unmarshal update response; cell:%v; err: %w", fraction.CellStr, err)
	}

	fmt.Printf(">>>>>>> 5: %+v\n", attributeMap)

	return nil
}

func (r *Repo) UpdateEntityType(ctx context.Context, PK, SK, entityType string) error {
	var err error
	update := expression.Set(expression.Name("EntityType"), expression.Value(entityType))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return fmt.Errorf("couldn't build expression for update fraction. err: %w", err)
	}

	pk, err := attributevalue.Marshal(PK)
	if err != nil {
		return fmt.Errorf("DANGER: failed to marshal pk; err: %w", err)
	}
	sk, err := attributevalue.Marshal(SK)
	if err != nil {
		return fmt.Errorf("DANGER: failed to marshal sk; err: %w", err)
	}
	itemKey := map[string]types.AttributeValue{"PK": pk, "SK": sk}

	_, err = r.core.DynamoDB().UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 table.Data(),
		Key:                       itemKey,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueNone,
	})
	if err != nil {
		return fmt.Errorf("couldn't update entity-type; PK: %q, SK: %q, err: %w", PK, SK, err)
	}

	return nil
}

func (r *Repo) UpdateFractionEntityType(ctx context.Context, fraction model.Fraction) error {
	update := expression.Set(expression.Name("EntityType"), expression.Value(fraction.EntityType))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return fmt.Errorf("couldn't build expression for update entity-type in fraction. err: %w", err)
	}

	if err := r.Update(ctx, fraction.GetKey(r.core), expr); err != nil {
		return fmt.Errorf("couldn't update entity-type for fraction. err: %w", err)
	}

	return nil
}
