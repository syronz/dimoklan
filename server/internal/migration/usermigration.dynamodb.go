package migration

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"dimoklan/consts"
	"dimoklan/types"
)

/*
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
*/

func (m Migration) AddUser() {
	user := types.User{
		ID:            "u#3224053",
		SK:            "u#3224053",
		Color:         "3131f5",
		Email:         "sabina.diako@gmail.com",
		Kingdom:       "Malanda",
		Password:      "6b53d67e399b703b38c58fa4c9e25438478ca0372b190abc2e34579e5e3cfa83",
		Language:      "en",
		Suspend:       false,
		SuspendReason: "",
		Freeze:        false,
		FreezeReason:  "",
		CreatedAt:     1709064739,
		UpdatedAt:     1709064739,
		EntityType:    "user",
	}

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("error in marshmap user; err: %v", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(consts.TableData),
	}

	if _, err = m.svc.PutItem(input); err != nil {
		log.Fatalf("error in creating user; err: %v", err)
	}

	fmt.Println("User added successfully")
}

/*
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
*/
