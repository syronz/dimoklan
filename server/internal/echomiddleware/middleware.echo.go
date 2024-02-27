package echomiddleware

import "dimoklan/internal/config"



type Middleware struct {
	core config.Core
}

func NewMiddleware(core config.Core) *Middleware {
	return &Middleware{
		core: core,
	}
}
