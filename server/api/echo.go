package api

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/storage"
	"dimoklan/types"
	"dimoklan/consts"
)

type Server struct {
	listenAddr string
	store      storage.Storage
	basStorage basstorage.BasStorage
}

func NewServer(cfg config.Core, store storage.Storage, basStorage basstorage.BasStorage) *Server {
	return &Server{
		listenAddr: cfg.GetPort(),
		store:      store,
		basStorage: basStorage,
	}
}

// Handler
func (s *Server) handleGetUserByID(c echo.Context) error {
	user := s.store.Get(10)

	return c.JSON(http.StatusOK, user)
}

func generateRandomColor() string {
	// Determine the number of bytes needed for a 6-character hex
	numBytes := 3

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// TODO: Handle the error (in a real application, you would handle errors appropriately)
		return "000000" // Default color in case of error
	}

	// Convert to hexadecimal string
	hexString := hex.EncodeToString(randomBytes)

	// Add a '#' prefix
	return hexString
}

func (s *Server) createUser(c echo.Context) error {
	var user types.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err,
		})
	}

	user.Code = uuid.New().String()
	user.Color = generateRandomColor()
	user.Status = consts.Active
	user.Reason = "new user"

	if err := s.basStorage.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, user)
}

var mu sync.Mutex
var index int

func (s *Server) Start() {
	e := echo.New()

	// rate limiter
	// Create a map to store rate limiters for each route
	//rateLimiters := make(map[string]*rate.Limiter)

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
	//e.POST("/users", s.createUser)

	e.Logger.Fatal(e.Start(s.listenAddr))
}
