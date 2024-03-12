package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	"dimoklan/api"
	"dimoklan/consts"
	"dimoklan/internal/config"
	"dimoklan/internal/echomiddleware"
	"dimoklan/repo"
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
	middleware := echomiddleware.NewMiddleware(s.core)
	e.Use(middleware.SetCustomContext)
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

	// Add middleware to log the user's IP address
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ip := req.Header.Get("X-Real-IP")
			if ip == "" {
				ip = req.Header.Get("X-Forwarded-For")
				if ip == "" {
					ip = req.RemoteAddr
				}
			}
			c.Set(consts.IP, ip)
			return next(c)
		}
	})

	storage := repo.NewRepo(s.core)
	cellService := service.NewCellService(s.core, storage)

	registerService := service.NewRegisterService(s.core, storage, cellService)
	registerAPI := api.NewRegisterAPI(s.core, registerService)

	authService := service.NewAuthService(s.core, storage)
	authAPI := api.NewAuthAPI(s.core, authService)

	fractionService := service.NewFractionService(s.core, storage)
	fractionAPI := api.NewFractionAPI(s.core, fractionService)

	marshalService := service.NewMarshalService(s.core, storage)
	marshalAPI := api.NewMarshalAPI(s.core, marshalService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.POST("/register", registerAPI.Create, defaultRateLimiter)
	e.GET("/register", registerAPI.Confirm, defaultRateLimiter)
	e.POST("/login", authAPI.Login)
	e.GET("/secure", func(c echo.Context) error {
		userID := c.Get("user_id")
		return c.String(http.StatusOK, fmt.Sprintf("Secure route for user_id: %v", userID))
	}, middleware.AuthMiddleware)
	e.GET("/fractions", fractionAPI.GetFraction, defaultRateLimiter, middleware.AuthMiddleware)
	e.GET("/marshals/:id", marshalAPI.GetMarshal, middleware.AuthMiddleware)
	e.POST("/marshals/:id/move", marshalAPI.MoveMarshal, middleware.AuthMiddleware)

	e.Logger.Fatal(e.Start(s.listenAddr))
}

func displayMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc: %v KB\tTotalAlloc: %v KB\tSys: %v MiB\tNumGC: %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024/1024, m.NumGC)
}

func main() {
	// Run the displayMemoryUsage function every 5 seconds
	// ticker := time.NewTicker(10 * time.Second)
	// defer ticker.Stop()

	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			displayMemoryUsage()
	// 		}
	// 	}
	// }()

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
