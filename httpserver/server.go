package httpserver

import (
	"encoding/json"
	"fmt"
	"dimoklan/points"
	"net/http"
)

var allPoints []points.Point // Variable to store all generated points

// GeneratePointsHandler handles the /generatePoints HTTP endpoint
func GeneratePointsHandler(w http.ResponseWriter, r *http.Request) {
	leftStr := r.URL.Query().Get("left")
	rightStr := r.URL.Query().Get("right")

	var left, right int
	// Parse the query parameters
	fmt.Sscanf(leftStr, "%d", &left)
	fmt.Sscanf(rightStr, "%d", &right)

	// Filter points based on the range
	filteredPoints := points.PointsInRange(left, right, allPoints)

	// Convert points to JSON
	responseJSON, err := json.Marshal(filteredPoints)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(responseJSON)
}

// StartServer starts the HTTP server
func StartServer(randomPoints []points.Point) {
	allPoints = randomPoints

	http.HandleFunc("/generatePoints", GeneratePointsHandler)
	http.Handle("/map/", http.FileServer(http.Dir("/home/diako/projects/dimoklan/client/")))
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("client"))))


	// Start the server on :8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

