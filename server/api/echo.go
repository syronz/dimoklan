package api

import (
	"fmt"
	"net/http"
	"sync"

	"dimoklan/internal/config"
	"dimoklan/service"
	"dimoklan/storage"
	"dimoklan/types"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

type Server struct {
	listenAddr  string
	store       storage.Storage
	userService *service.UserService
}

func NewServer(core config.Core, store storage.Storage, userService *service.UserService) *Server {
	return &Server{
		listenAddr:  core.GetPort(),
		store:       store,
		userService: userService,
	}
}

// Handler
func (s *Server) handleGetUserByID(c echo.Context) error {
	user := s.store.Get(10)

	return c.JSON(http.StatusOK, user)
}

func (s *Server) createUser(c echo.Context) error {
	var user types.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err,
		})
	}

	var err error
	if user, err = s.userService.Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

var (
	mu    sync.Mutex
	index int
)

func (s *Server) Start() {
	e := echo.New()

	// rate limiter
	// Create a map to store rate limiters for each route
	// rateLimiters := make(map[string]*rate.Limiter)

	limiter := rate.NewLimiter(rate.Limit(80000), 200000)
	// Middleware function to enforce rate limiting based on the route
	createUserRateLimiter := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mu.Lock()
			index++
			if index%10000 == 0 {
				fmt.Println(">>i", index)
			}
			mu.Unlock()
			if !limiter.Allow() {
				return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "Rate limit exceeded"})
			}

			return next(c)
		}
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/users", s.handleGetUserByID)
	e.POST("/users", s.createUser, createUserRateLimiter)
	// e.POST("/users", s.createUser)

	e.Logger.Fatal(e.Start(s.listenAddr))
}
