package mapmigration

import (
	"fmt"
	"log"

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

func (m Migration) CreateCellTable() {
	// Define table name
	tableName := "cells"

	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
	if err != nil {
		// Table doesn't exist, create it
		fmt.Println("Table doesn't exist. Creating table...")

		// Define table schema
		createTableInput := &dynamodb.CreateTableInput{
			TableName: aws.String(tableName),
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("cell"),
					AttributeType: aws.String("S"), // Assuming "cell" is a string
				},
				{
					AttributeName: aws.String("user_id"),
					AttributeType: aws.String("N"), // "user_id" is a number
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("cell"),
					KeyType:       aws.String("HASH"), // Partition key
				},
				{
					AttributeName: aws.String("user_id"),
					KeyType:       aws.String("RANGE"), // Sort key
				},
			},
			BillingMode: aws.String("PAY_PER_REQUEST"),
		}

		_, err := m.svc.CreateTable(createTableInput)
		if err != nil {
			log.Fatalf("Error creating table:", err)
		}

		fmt.Println("Table created successfully:", tableName)
	} else {
		fmt.Println("Table already exists:", tableName)
	}
}
