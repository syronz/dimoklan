package model

import "dimoklan/model/localtype"

type Fraction struct {
	Fraction      string    `json:"fraction"`
	Cells         []Cell    `json:"cells"`
	MovingMarshal []Marshal `json:"moving_marshals"`
	FixedMarshal  []Marshal `json:"fixed_marshals"`
}

type FractionRepo struct {
	Fraction   string `dynamodbav:"PK"`
	SK         string `dynamodbav:"SK"`
	EntityType string
	UserID     string
	Building   string
	Score      int
	Cell       localtype.CELL
	Name       string
	Army       int
	Star       int
	Speed      float64
	Attack     float64
	Face       string
	CreatedAt  int64
	UpdatedAt  int64
}
