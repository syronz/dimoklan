package basmigration

import (
	"fmt"
	"log"

	"dimoklan/consts"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (m Migration) CreateDataTable() {
	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableData),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
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

	_, err = m.svc.CreateTable(createTableInput)
	if err != nil {
		log.Fatalf("Error creating table: %v; %v", consts.TableData, err)
	}

	fmt.Println("Table created successfully:", consts.TableData)

	// Wait for table creation to complete (Optional)
	// Note: This step is optional and depends on your use case
	err = m.svc.WaitUntilTableExists(&dynamodb.DescribeTableInput{
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

	_, err = m.svc.UpdateTimeToLive(updateInput)
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
	_, err := m.svc.DeleteTable(input)
	if err != nil {
		log.Printf("Error deleting table: %v; %v", consts.TableData, err)
		return
	}
}
