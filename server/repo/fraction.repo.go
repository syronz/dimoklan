package repo

import (
	_ "embed"

	"dimoklan/model"
)

func (r *Repo) GetFractions(coordinates []string) ([]model.Fraction, error) {
	fractions := make([]model.Fraction, len(coordinates))

	/*
		for _, coordinate := range coordinates {
			coordinate = consts.ParFraction + coordinate
			params := &dynamodb.QueryInput{
				TableName:              aws.String(consts.TableData),
				KeyConditionExpression: aws.String("PK = :fraction"),
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":fraction": {
						S: aws.String(coordinate),
					},
				},
			}

			resp, err := r.core.DynamoDB().Query(params)
			if err != nil {
				return fractions, fmt.Errorf("error in getting fractions; err: %w", err)
			}

			fmt.Println(">>>>--", resp, params)

			// err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, fractions[i].Cells)

		}
	*/

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
