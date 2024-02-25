package basstorage

import "dimoklan/internal/config"

type BasMysql struct {
	core config.Core
}

func NewBasMysql(core config.Core) *BasMysql {
	return &BasMysql{
		core: core,
	}
}
