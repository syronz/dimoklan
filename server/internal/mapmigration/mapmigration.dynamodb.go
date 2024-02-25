package mapmigration

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Define table name
const mapTable = "map"

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
		TableName: aws.String(mapTable),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
	if err != nil {
		// Table doesn't exist, create it
		fmt.Println("Table doesn't exist. Creating table...")

		createTableInput := &dynamodb.CreateTableInput{
			TableName: aws.String(mapTable),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("fraction"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("cell"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("user_id"),
					AttributeType: aws.String("N"),
				},
				{
					AttributeName: aws.String("last_update"),
					AttributeType: aws.String("N"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("fraction"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("cell"),
					KeyType:       aws.String("RANGE"),
				},
			},
			BillingMode: aws.String("PAY_PER_REQUEST"),
			GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
				{
					IndexName: aws.String("user_id_index"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("user_id"),
							KeyType:       aws.String("HASH"),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String("ALL"),
					},
				},
				{
					IndexName: aws.String("cell_index"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("cell"),
							KeyType:       aws.String("HASH"),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String("ALL"),
					},
				},
				{
					IndexName: aws.String("last_update_index"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("fraction"),
							KeyType:       aws.String("HASH"),
						},
						{
							AttributeName: aws.String("last_update"),
							KeyType:       aws.String("RANGE"),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String("ALL"),
						// ProjectionType: aws.String("INCLUDE"),
						// NonKeyAttributes: []*string{
						// 	aws.String("cell"),
						// },
					},
				},
			},
		}

		_, err := m.svc.CreateTable(createTableInput)
		if err != nil {
			log.Fatalf("Error creating table:", err)
		}

		fmt.Println("Table created successfully:", mapTable)
	} else {
		fmt.Println("Table already exists:", mapTable)
	}
}

func (m Migration) DeleteMapTable() {

	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(mapTable),
	}

	// Delete the table
	_, err := m.svc.DeleteTable(input)
	if err != nil {
		log.Fatalf("Error deleting table: %v; %v", mapTable, err)
		return
	}
}
