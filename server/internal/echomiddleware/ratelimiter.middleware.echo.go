package echomiddleware

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

var (
	mu    sync.Mutex
	index int
)

func (m *Middleware) DefaultRateLimiter(a, b int) func(echo.HandlerFunc) echo.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(a), b)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// rate limiter
		// Create a map to store rate limiters for each route
		// rateLimiters := make(map[string]*rate.Limiter)

		// limiter := rate.NewLimiter(rate.Limit(80000), 200000)
		// defaultRateLimiter := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mu.Lock()
			index++
			if index%1 == 0 {
				fmt.Println(">>i", index)
			}
			mu.Unlock()
			if !limiter.Allow() {
				return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "Rate limit exceeded"})
			}

			return next(c)
		}
	}
}
