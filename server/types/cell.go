package types

import "time"

type Cell struct {
	X         int       `json:"x"`
	Y         int       `json:"y"`
	UserID    int       `json:"user_id"`
	Building  string    `json:"building"`
	Score     int       `json:"score"`
	UpdatedAt time.Time `json:"updated_at"`
}

func validateCell(c *Cell) bool { return true }