package mapmigration

import (
	"fmt"
	"log"

	"dimoklan/consts"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Migration struct {
	svc *dynamodb.DynamoDB
}

func New(region, endpoint string) Migration {
	// Create a session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),   // Change the region according to your setup
		Endpoint: aws.String(endpoint), // Local DynamoDB endpoint
	}))

	// Create a DynamoDB service client
	return Migration{
		svc: dynamodb.New(sess),
	}
}

func (m Migration) CreateMapTable() {

	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableMap),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
	if err != nil {
		fmt.Println("Table doesn't exist. Creating table...", consts.TableMap)

		createTableInput := &dynamodb.CreateTableInput{
			TableName: aws.String(consts.TableMap),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("fraction"),
					AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
				},
				{
					AttributeName: aws.String("cell"),
					AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
				},
				{
					AttributeName: aws.String("user_id"),
					AttributeType: aws.String(dynamodb.ScalarAttributeTypeN),
				},
				{
					AttributeName: aws.String("last_update"),
					AttributeType: aws.String(dynamodb.ScalarAttributeTypeN),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("fraction"),
					KeyType:       aws.String(dynamodb.KeyTypeHash),
				},
				{
					AttributeName: aws.String("cell"),
					KeyType:       aws.String(dynamodb.KeyTypeRange),
				},
			},
			BillingMode: aws.String("PAY_PER_REQUEST"),
			GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
				{
					IndexName: aws.String("user_id_index"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("user_id"),
							KeyType:       aws.String(dynamodb.KeyTypeHash),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
					},
				},
				{
					IndexName: aws.String("cell_index"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("cell"),
							KeyType:       aws.String(dynamodb.KeyTypeHash),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
					},
				},
				// {
				// 	IndexName: aws.String("last_update_index"),
				// 	KeySchema: []*dynamodb.KeySchemaElement{
				// 		{
				// 			AttributeName: aws.String("fraction"),
				// 			KeyType:       aws.String("HASH"),
				// 		},
				// 		{
				// 			AttributeName: aws.String("last_update"),
				// 			KeyType:       aws.String("RANGE"),
				// 		},
				// 	},
				// 	Projection: &dynamodb.Projection{
				// 		ProjectionType: aws.String("INCLUDE"),
				// 		NonKeyAttributes: []*string{
				// 			aws.String("cell"),
				// 		},
				// 	},
				// },
			},
		}

		_, err := m.svc.CreateTable(createTableInput)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}

		fmt.Println("Table created successfully:", consts.TableMap)
	} else {
		fmt.Println("Table already exists:", consts.TableMap)
	}
}

func (m Migration) DeleteMapTable() {

	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(consts.TableMap),
	}

	// Delete the table
	_, err := m.svc.DeleteTable(input)
	if err != nil {
		log.Fatalf("Error deleting table: %v; %v", consts.TableMap, err)
		return
	}
}
