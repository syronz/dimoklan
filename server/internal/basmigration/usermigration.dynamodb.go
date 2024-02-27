package basmigration

import (
	"fmt"
	"log"

	"dimoklan/consts"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (m Migration) CreateUserTable() {
	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableUser),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
	if err == nil {
		log.Fatalf("table already exist: %v, err:%v", consts.TableUser, err)
	}

	fmt.Println("Table doesn't exist. Creating table...", consts.TableUser)
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(consts.TableUser),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeN),
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String(consts.IndexEmail),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("email"),
						KeyType:       aws.String(dynamodb.KeyTypeHash),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
				},
			},
		},
	}

	_, err = m.svc.CreateTable(createTableInput)
	if err != nil {
		log.Fatalf("Error creating table: %v; %v", consts.TableUser, err)
	}

	fmt.Println("Table created successfully:", consts.TableUser)
}

func (m Migration) DeleteUserTable() {
	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(consts.TableUser),
	}

	// Delete the table
	_, err := m.svc.DeleteTable(input)
	if err != nil {
		log.Printf("Error deleting table: %v; %v", consts.TableUser, err)
		return
	}
}
