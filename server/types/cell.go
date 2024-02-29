package types

import "dimoklan/types/localtype"

type Cell struct {
	Fraction string         `json:"fraction" dynamodbav:"PK"`
	Cell     localtype.CELL `json:"cell" dynamodbav:"SK"`
	UserID   string         `json:"user_id" dynamodbav:"UserID"`
	Building string         `json:"building" dynamodbav:"Building"`
	Score    int            `json:"score" dynamodbav:"Score"`

	// internal attributes
	UpdatedAt  int64  `json:"last_update" dynamodbav:"UpdatedAt"`
	EntityType string `json:"-" dynamodbav:"EntityType"`
}

func validateCell(c *Cell) bool { return true }
