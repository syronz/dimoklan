package model

import (
	"time"

	"dimoklan/consts/entity"
	"dimoklan/model/localtype"
)

type Marshal struct {
	UserID    string         `json:"user_id"`
	ID        string         `json:"id"`
	Cell      localtype.CELL `json:"cell"`
	Name      string         `json:"name"`
	Army      int            `json:"army"`
	Star      int            `json:"star"`
	Speed     float64        `json:"speed"`
	Attack    float64        `json:"attack"`
	Face      string         `json:"face"`
	CreatedAt time.Time      `json:"created_at"`
}

type MarshalRepo struct {
	PK         string  `dynamodbav:"PK"` // userID
	SK         string  `dynamodbav:"SK"` // marshalID
	Cell       string  `dynamodbav:"Cell"`
	Name       string  `dynamodbav:"Name"`
	Army       int     `dynamodbav:"Army"`
	Star       int     `dynamodbav:"Star"`
	Speed      float64 `dynamodbav:"Speed"`
	Attack     float64 `dynamodbav:"Attack"`
	Face       string  `dynamodbav:"Face"`
	CreatedAt  int64   `dynamodbav:"CreatedAt"`
	EntityType string  `dynamodbav:"EntityType"`
}

func (m *Marshal) ToRepo() MarshalRepo {
	return MarshalRepo{
		PK:         m.UserID,
		SK:         m.ID,
		Cell:       m.Cell.ToString(),
		Name:       m.Name,
		Army:       m.Army,
		Star:       m.Star,
		Speed:      m.Speed,
		Attack:     m.Attack,
		Face:       m.Face,
		CreatedAt:  m.CreatedAt.Unix(),
		EntityType: entity.Marshal,
	}
}

func (m *MarshalRepo) ToAPI() Marshal {
	return Marshal{
		UserID:    m.PK,
		ID:        m.SK,
		Cell:      localtype.ToCell(m.Cell),
		Name:      m.Name,
		Army:      m.Army,
		Star:      m.Star,
		Speed:     m.Speed,
		Attack:    m.Attack,
		Face:      m.Face,
		CreatedAt: time.Unix(m.CreatedAt, 0),
	}
}

func validateMarshal(c *Marshal) bool { return true }
