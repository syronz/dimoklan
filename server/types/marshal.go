package types

type Marshal struct {
	UserID     string  `json:"user_id" dynamodbav:"PK"`
	ID         string  `json:"id" dynamodbav:"SK"`
	Cell       CELL    `json:"cell" dynamodbav:"Cell"`
	Name       string  `json:"name" dynamodbav:"Name"`
	Army       int     `json:"army" dynamodbav:"Army"`
	Star       int     `json:"star" dynamodbav:"Star"`
	Speed      float64 `json:"speed" dynamodbav:"Speed"`
	Attack     float64 `json:"attack" dynamodbav:"Attack"`
	Face       string  `json:"face" dynamodbav:"Face"`
	CreatedAt  int64   `json:"created_at" dynamodbav:"CreatedAt"`
	EntityType string  `json:"-" dynamodbav:"EntityType"`
}

func validateMarshal(c *Marshal) bool { return true }
