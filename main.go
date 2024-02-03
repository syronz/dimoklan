package main

import (
	"fmt"

	"dimoklan/httpserver"
	"dimoklan/points"
)

// const numRandomPoints = 2_000_000 // Adjust the number of random points as needed
const numRandomPoints = 200 // Adjust the number of random points as needed

func main() {
	// Generate random points at the start
	// randomPoints := points.GenerateRandomPoints(numRandomPoints, 10_000_000)
	randomPoints := points.GenerateRandomPoints(numRandomPoints, 1000)
	// Start the server
	fmt.Println("Server is running on :8080...")
	fmt.Print("Server is running on :8080...")
	httpserver.StartServer(randomPoints)
}
