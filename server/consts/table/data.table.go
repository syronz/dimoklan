package table

import (
	"github.com/aws/aws-sdk-go-v2/aws"

	"dimoklan/consts"
)

func Data() *string {
	return aws.String(consts.TableData)
}
