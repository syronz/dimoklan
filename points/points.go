package points

import (
	"math/rand"
	"time"
)

// Point represents a 2D point with X and Y coordinates
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// GenerateRandomPoints generates a specified count of random points
func GenerateRandomPoints(count int, maxCoordinate int) []Point {
	rand.Seed(time.Now().UnixNano())
	var points []Point
	for i := 0; i < count; i++ {
		point := Point{
			X: rand.Intn(maxCoordinate), // Adjust the range as needed
			Y: rand.Intn(maxCoordinate),
		}
		points = append(points, point)
	}

	return points
}


// PointsInRange filters points based on the specified range
func PointsInRange(left, right int, points []Point) []Point {
	var result []Point
	for _, point := range points {
		if point.X >= left && point.X <= right {
			result = append(result, point)
		}
	}
	return result
}
