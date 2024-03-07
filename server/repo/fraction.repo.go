package repo

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

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

/*
// create the api params

	// dump the response data
	fmt.Println(resp)

	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var movies []model.Movie
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &movies)

	// print the response data
	for _, m := range movies {
		fmt.Printf("Movie: '%s' (%d)\n", m.Title, m.Year)
	}
*/
