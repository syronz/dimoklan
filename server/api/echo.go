package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"dimoklan/internal/config"
	"dimoklan/service"
	"dimoklan/storage"
	"dimoklan/types"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	// images

	"image/color"
	"strconv"

	"github.com/fogleman/gg"
)

type Server struct {
	listenAddr  string
	store       storage.Storage
	userService *service.UserService
	cellService *service.CellService
}

func NewServer(
	core config.Core,
	store storage.Storage,
	userService *service.UserService,
	cellService *service.CellService,
) *Server {
	return &Server{
		listenAddr:  core.GetPort(),
		store:       store,
		userService: userService,
		cellService: cellService,
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

func generateImage(c echo.Context) error {
	// Parse query parameters
	width, err := strconv.Atoi(c.QueryParam("width"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid width parameter")
	}

	height, err := strconv.Atoi(c.QueryParam("height"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid height parameter")
	}

	// Create an image
	img := gg.NewContext(width, height)
	img.SetColor(color.RGBA{255, 255, 255, 255})
	img.Clear()

	// Draw something on the image (e.g., a rectangle)
	img.SetColor(color.RGBA{0, 0, 255, 255})
	img.DrawRectangle(10, 10, float64(width-20), float64(height-20))
	img.Fill()

	// Save the image as a response
	// var imageBuffer []byte
	if err := gg.SavePNG("./image.png", img.Image()); err != nil {
		return err
	}

	imageBytes, err := ioutil.ReadFile("image.png")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error reading image file")
	}

	return c.Blob(http.StatusOK, "image/png", imageBytes)
}

func (s *Server) GenerateImage(c echo.Context) error {
	start := types.Point{X: 0, Y: 0}
	stop := types.Point{X: 1000, Y: 1000}

	pixelData, err := s.cellService.GetMap(start, stop)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
	}

	// Example pixel data (replace with your own)
	// pixelData := [][]string{
	// 	{"#FF0000", "#00FF00", "#0000FF"},
	// 	{"#FFFF00", "#FF00FF", "#00FFFF"},
	// 	{"#000000", "#FFFFFF", "#888888"},
	// }

	// Calculate image dimensions based on pixelData
	rows := len(pixelData)
	cols := len(pixelData[0])
	cellSize := 3
	width := cols * cellSize
	height := rows * cellSize

	// Create an image
	img := gg.NewContext(width, height)
	img.SetColor(color.White)
	img.Clear()

	// Draw pixels based on pixelData
	for i, row := range pixelData {
		for j, hexCode := range row {
			if hexCode == "" {
				hexCode = "EEEEEE"
			}
			hexCode = "#" + hexCode
			color := parseHexCode(hexCode)
			img.SetColor(color)
			img.DrawRectangle(float64(j*cellSize), float64(i*cellSize), float64(cellSize), float64(cellSize))
			img.Fill()
		}
	}

	writer := c.Response().Writer
	img.EncodePNG(writer)
	return nil
}

func parseHexCode(hexCode string) color.RGBA {
	c := color.RGBA{}
	c.A = 0xFF
	if len(hexCode) == 7 && hexCode[0] == '#' {
		if r, err := strconv.ParseUint(hexCode[1:3], 16, 8); err == nil {
			c.R = uint8(r)
		}
		if g, err := strconv.ParseUint(hexCode[3:5], 16, 8); err == nil {
			c.G = uint8(g)
		}
		if b, err := strconv.ParseUint(hexCode[5:7], 16, 8); err == nil {
			c.B = uint8(b)
		}
	}
	return c
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
	e.GET("/generate-image", s.GenerateImage)

	e.Logger.Fatal(e.Start(s.listenAddr))
}
