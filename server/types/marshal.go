package types

import "time"

type Marshal struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Star      string    `json:"star"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func validateMarshal(c *Marshal) bool { return true }
