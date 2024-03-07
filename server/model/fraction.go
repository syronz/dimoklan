package model

import "dimoklan/model/localtype"

type Fraction struct {
	Fraction   string         `json:"fraction,omitempty" dynamodbav:"PK"`
	Cell       localtype.CELL `json:"cell,omitempty" dynamodbav:"SK"`
	EntityType string         `json:"entity_type,omitempty" dynamodbav:"EntityType"`
	UserID     string         `json:"user_id,omitempty" dynamodbav:"UserID"`
	Building   string         `json:"building,omitempty" dynamodbav:"Building"`
	Score      int            `json:"score,omitempty" dynamodbav:"Score"`
	Name       string         `json:"name,omitempty" dynamodbav:"Name"`
	Army       int            `json:"army,omitempty" dynamodbav:"Army"`
	Star       int            `json:"start,omitempty" dynamodbav:"Start"`
	Speed      float64        `json:"speed,omitempty" dynamodbav:"Speed"`
	Attack     float64        `json:"attack,omitempty" dynamodbav:"Attack"`
	Face       string         `json:"face,omitempty" dynamodbav:"Face"`
	CreatedAt  int64          `json:"created_at,omitempty" dynamodbav:"CreatedAt"`
	UpdatedAt  int64          `json:"updated_at,omitempty" dynamodbav:"UpdatedAt"`
}
