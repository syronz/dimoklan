package table

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

func Data() *string {
	return aws.String("data")
}
