package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	"dimoklan/domain/basic/basapi"
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/service"
)

var configFilePath = flag.String("cfg", "", "config file path")
var (
	mu    sync.Mutex
	index int
)

type Server struct {
	core       config.Core
	listenAddr string
}

func newServer(
	core config.Core,
) *Server {
	return &Server{
		core:       core,
		listenAddr: core.GetPort(),
	}
}

func (s *Server) start() {
	e := echo.New()

	// rate limiter
	// Create a map to store rate limiters for each route
	// rateLimiters := make(map[string]*rate.Limiter)

	limiter := rate.NewLimiter(rate.Limit(80000), 200000)
	// Middleware function to enforce rate limiting based on the route
	defaultRateLimiter := func(next echo.HandlerFunc) echo.HandlerFunc {
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


	basStorage := basstorage.NewBasDynamoDB(s.core)

	userService := service.NewUserService(s.core, basStorage)
	userAPI := basapi.NewBasUserAPI(s.core, userService)

	registerService := service.NewRegisterService(s.core, basStorage)
	registerAPI := basapi.NewBasRegisterAPI(s.core, registerService)


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.POST("/users", userAPI.CreateUser, defaultRateLimiter)
	e.POST("/register", registerAPI.CreateRegister, defaultRateLimiter)

	e.Logger.Fatal(e.Start(s.listenAddr))
}

func main() {
	flag.Parse()

	if *configFilePath == "" {
		log.Fatal("cfg is required")
	}

	core, err := config.GetCore(*configFilePath)
	if err != nil {
		log.Fatalf("error in loading core; %v", err)
	}

	core.Info("starting server: " + time.Now().String())

	server := newServer(core)
	server.start()
}
