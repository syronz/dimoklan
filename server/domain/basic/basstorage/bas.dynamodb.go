package basstorage

import (
	"github.com/aws/aws-sdk-go/aws"

	"dimoklan/consts"
	"dimoklan/internal/config"
)

type BasDynamoDB struct {
	core          config.Core
	registerTable *string
}

func NewBasDynamoDB(core config.Core) *BasDynamoDB {
	return &BasDynamoDB{
		core: core,
		registerTable: aws.String(consts.TableRegister),
	}
}
