package echomiddleware

import (
	"context"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

func (c *CustomContext) Convert() context.Context {
	req := c.Request()
	ctx := context.WithValue(req.Context(), "user_id", c.Get("user_id"))
	return ctx
}

func (m *Middleware) SetCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		return next(cc)
	}
}
