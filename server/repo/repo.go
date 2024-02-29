package repo

import (
	"dimoklan/internal/config"
)

type Repo struct {
	core          config.Core
	registerTable *string
}

func NewRepo(core config.Core) *Repo {
	return &Repo{
		core: core,
	}
}
