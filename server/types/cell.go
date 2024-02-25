package types

type Cell struct {
	Fraction string `json:"fraction"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	UserID   int    `json:"user_id"`
	Building string `json:"building"`
	Score    int    `json:"score"`

	// internal attributes
	Cell       string `json:"cell"`
	LastUpdate int64  `json:"last_update"`
}

func validateCell(c *Cell) bool { return true }
