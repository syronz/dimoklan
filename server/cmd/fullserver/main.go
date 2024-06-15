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

	middleware := echomiddleware.NewMiddleware(s.core)

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

	e.Use(middleware.DefaultRateLimiter(1, 2))

	e.POST("/register", registerAPI.Create)
	e.GET("/register", registerAPI.Confirm)
	e.POST("/login", authAPI.Login)

	play := e.Group("/play")
	{
		// IMPORTANT to have these two middleware to have access to user_id
		play.Use(middleware.AuthMiddleware)
		play.Use(middleware.ContextGenerator)

		play.GET("/secure", func(c echo.Context) error {
			userID := c.Get("user_id")
			return c.String(http.StatusOK, fmt.Sprintf("Secure route for user_id: %v", userID))
		})
		play.GET("/fractions", fractionAPI.GetFraction)
		play.GET("/marshals/:id", marshalAPI.GetMarshal)
		play.POST("/marshals/:id/move", marshalAPI.MoveMarshal)
	}

	e.Logger.Fatal(e.Start(s.listenAddr))
}

func displayMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc: %.2f MB\tTotalAlloc: %.2f MB\tSys: %.2f MiB\tNumGC: %v\n",
		float64(m.Alloc)/1024/1024, float64(m.TotalAlloc)/1024/1024, float64(m.Sys)/1024/1024, m.NumGC)
}

func main() {
	// Run the displayMemoryUsage function every 5 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				// displayMemoryUsage()
			}
		}
	}()

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
