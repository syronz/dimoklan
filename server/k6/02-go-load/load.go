package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	startTime := time.Now()
	// Send POST request
	url := "http://127.0.0.1:3000/users"
	var i int
	for i = 0; i < 200002; i++ {
		time.Sleep(100 * time.Microsecond)
		go func() {
			// Generate random data
			name := getRandomString(8)
			username := getRandomString(8)
			password := getRandomString(12)
			createdAt := time.Now().UTC()
			updatedAt := time.Now().UTC()

			// Create User object
			user := User{
				Name:      name,
				Username:  username,
				Password:  password,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			}

			// Convert User object to JSON
			jsonData, err := json.Marshal(user)
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error sending POST request:", err)
				return
			}
			if resp.StatusCode != 200 {
				fmt.Println("status code is ", resp.StatusCode)
			}
			defer resp.Body.Close()
		}()
	}

	// Print the response status
	fmt.Printf("%v requests sent\n", i)
	fmt.Println("duration:", time.Since(startTime))
	time.Sleep(10 * time.Second)
}

func getRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
