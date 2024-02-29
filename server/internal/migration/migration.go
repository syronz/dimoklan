package migration

import (
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
