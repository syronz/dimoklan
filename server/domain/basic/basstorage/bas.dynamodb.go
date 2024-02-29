package basstorage

import (
	"dimoklan/internal/config"
)

type BasDynamoDB struct {
	core          config.Core
	registerTable *string
}

func NewBasDynamoDB(core config.Core) *BasDynamoDB {
	return &BasDynamoDB{
		core: core,
	}
}
