package migration

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"dimoklan/consts"
)

/*
func (m Migration) CreateDataTable() {
	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableData),
	}
	_, err := m.client.DescribeTable(describeTableInput)
	if err == nil {
		log.Fatalf("table already exist: %v, err:%v", consts.TableData, err)
	}

	fmt.Println("Table doesn't exist. Creating table...", consts.TableData)
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(consts.TableData),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
	}

	_, err = m.client.CreateTable(createTableInput)
	if err != nil {
		log.Fatalf("Error creating table: %v; %v", consts.TableData, err)
	}

	fmt.Println("Table created successfully:", consts.TableData)

	// Wait for table creation to complete (Optional)
	// Note: This step is optional and depends on your use case
	err = m.client.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableData),
	})
	if err != nil {
		log.Fatalf("Error waiting for table: %v; %v", consts.TableData, err)
	}

	// Enable TTL on the table
	ttlSpecification := &dynamodb.TimeToLiveSpecification{
		AttributeName: aws.String("ttl"),
		Enabled:       aws.Bool(true),
	}

	updateInput := &dynamodb.UpdateTimeToLiveInput{
		TableName:               aws.String(consts.TableData),
		TimeToLiveSpecification: ttlSpecification,
	}

	_, err = m.client.UpdateTimeToLive(updateInput)
	if err != nil {
		log.Fatalf("Error updating ttl in table: %v; %v", consts.TableData, err)
	}

	fmt.Println("TTL enabled successfully")

}

func (m Migration) DeleteDataTable() {

	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(consts.TableData),
	}

	// Delete the table
	_, err := m.client.DeleteTable(input)
	if err != nil {
		log.Printf("Error deleting table: %v; %v", consts.TableData, err)
		return
	}
}
*/

// TableExists determines whether a DynamoDB table exists.
func (m Migration) TableExists() (bool, error) {
	exists := true
	_, err := m.client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(consts.TableData)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", consts.TableData)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %q. Here's why: %v\n", consts.TableData, err)
		}
		exists = false
	}
	return exists, err
}

func (m Migration) CreateDataTable() {
	isTableExist, err := m.TableExists()
	if isTableExist {
		log.Fatalf("table already exist: %v", consts.TableData)
	}
	if err != nil {
		log.Fatalf("error in checking table existance: %v, err:%v", consts.TableData, err)
	}

	fmt.Println("Table doesn't exist. Creating table...", consts.TableData)
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(consts.TableData),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	}

	_, err = m.client.CreateTable(context.TODO(), createTableInput)
	if err != nil {
		log.Fatalf("Error creating table: %v; %v", consts.TableData, err)
	}

	// TODO: add ttl column

}

func (m Migration) DeleteDataTable() {
	_, err := m.client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(consts.TableData)})
	if err != nil {
		log.Fatalf("Couldn't delete table %v. Here's why: %v\n", consts.TableData, err)
	}

	/*
		input := &dynamodb.DeleteTableInput{
			TableName: aws.String(consts.TableData),
		}

		// Delete the table
		_, err := m.client.DeleteTable(input)
		if err != nil {
			log.Printf("Error deleting table: %v; %v", consts.TableData, err)
			return
		}
	*/
}
