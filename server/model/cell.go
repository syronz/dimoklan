package model

import (
	"time"

	"dimoklan/consts/entity"
	"dimoklan/model/localtype"
)

type Cell struct {
	Fraction  string         `json:"fraction"`
	Cell      localtype.CELL `json:"cell"`
	UserID    string         `json:"user_id"`
	Building  string         `json:"building"`
	Score     int            `json:"score"`
	UpdatedAt time.Time      `json:"last_update"`
}

type CellRepo struct {
	PK         string         `dynamodbav:"PK"`
	SK         localtype.CELL `dynamodbav:"SK"`
	UserID     string         `dynamodbav:"UserID"`
	Building   string         `dynamodbav:"Building"`
	Score      int            `dynamodbav:"Score"`
	UpdatedAt  int64          `dynamodbav:"UpdatedAt"`
	EntityType string         `dynamodbav:"EntityType"`
}

func (c *Cell) ToRepo() CellRepo {
	return CellRepo{
		PK:         c.Fraction,
		SK:         c.Cell,
		UserID:     c.UserID,
		Building:   c.Building,
		Score:      c.Score,
		UpdatedAt:  c.UpdatedAt.Unix(),
		EntityType: entity.Cell,
	}
}

func (c *CellRepo) ToAPI() Cell {
	return Cell{
		Fraction:  c.PK,
		Cell:      c.SK,
		UserID:    c.UserID,
		Building:  c.Building,
		Score:     c.Score,
		UpdatedAt: time.Unix(c.UpdatedAt, 0),
	}
}

func validateCell(c *Cell) bool { return true }
