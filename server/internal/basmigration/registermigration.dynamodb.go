package basmigration

import (
	"fmt"
	"log"

	"dimoklan/consts"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (m Migration) CreateRegisterTable() {
	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableRegister),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
	if err == nil {
		log.Fatalf("register table already exist: %v", err)
	}

	fmt.Println("Table doesn't exist. Creating table...", consts.TableRegister)
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(consts.TableRegister),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("hash"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ttl"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("hash"),
				KeyType:       aws.String(dynamodb.KeyTypeHash),
			},
			{
				AttributeName: aws.String("ttl"),
				KeyType:       aws.String(dynamodb.KeyTypeRange),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
	}

	_, err = m.svc.CreateTable(createTableInput)
	if err != nil {
		log.Fatalf("Error creating table: %v; %v", consts.TableRegister, err)
	}

	fmt.Println("Table created successfully:", consts.TableRegister)

	// Wait for table creation to complete (Optional)
	// Note: This step is optional and depends on your use case
	err = m.svc.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: aws.String(consts.TableRegister),
	})
	if err != nil {
		log.Fatalf("Error waiting for table: %v; %v", consts.TableRegister, err)
	}

	// Enable TTL on the table
	ttlSpecification := &dynamodb.TimeToLiveSpecification{
		AttributeName: aws.String("ttl"),
		Enabled:       aws.Bool(true),
	}

	updateInput := &dynamodb.UpdateTimeToLiveInput{
		TableName:               aws.String(consts.TableRegister),
		TimeToLiveSpecification: ttlSpecification,
	}

	_, err = m.svc.UpdateTimeToLive(updateInput)
	if err != nil {
		log.Fatalf("Error updating ttl in table: %v; %v", consts.TableRegister, err)
	}

	fmt.Println("TTL enabled successfully")

}

func (m Migration) DeleteRegisterTable() {

	input := &dynamodb.DeleteTableInput{
		TableName: aws.String(consts.TableRegister),
	}

	// Delete the table
	_, err := m.svc.DeleteTable(input)
	if err != nil {
		log.Fatalf("Error deleting table: %v; %v", consts.TableRegister, err)
		return
	}
}
