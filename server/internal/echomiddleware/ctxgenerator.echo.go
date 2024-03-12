package echomiddleware

import (
	"context"

	"dimoklan/consts/ctxkey"

	"github.com/labstack/echo/v4"
)

func (m *Middleware) ContextGenerator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the existing request
		req := c.Request()

		// Update the context
		newContext := context.WithValue(req.Context(), ctxkey.UserID, c.Get("user_id"))

		// Create a new request with the updated context
		newReq := req.WithContext(newContext)

		// Update the context in the Echo context
		c.SetRequest(newReq)

		// Call the next middleware or handler
		return next(c)
	}
}
