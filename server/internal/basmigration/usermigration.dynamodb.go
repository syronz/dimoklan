package basmigration

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const usersTable = "users"

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
	// Check if the table exists
	describeTableInput := &dynamodb.DescribeTableInput{
		TableName: aws.String(usersTable),
	}
	_, err := m.svc.DescribeTable(describeTableInput)
	if err != nil {
		log.Fatalf("users table already exist: %v", err)
	}

	fmt.Println("Table doesn't exist. Creating table...", usersTable)
	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(usersTable),
	}

	_ = createTableInput

}
