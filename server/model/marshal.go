package model

import (
	"time"

	"dimoklan/consts/entity"
	"dimoklan/consts/hashtag"
	"dimoklan/model/localtype"
)

type Marshal struct {
	UserID    string         `json:"user_id" dynamodbav:"PK"`
	ID        string         `json:"id" dynamodbav:"SK"`
	Cell      localtype.CELL `json:"cell" dynamodbav:"Cell"`
	Name      string         `json:"name" dynamodbav:"Name"`
	Army      int            `json:"army" dynamodbav:"Army"`
	Star      int            `json:"star" dynamodbav:"Star"`
	Speed     float64        `json:"speed" dynamodbav:"Speed"`
	Attack    float64        `json:"attack" dynamodbav:"Attack"`
	Face      string         `json:"face" dynamodbav:"Face"`
	CreatedAt time.Time      `json:"created_at" dynamodbav:"CreatedAt"`
}

type MarshalRepo struct {
	UserID     string         `dynamodbav:"PK"`
	ID         string         `dynamodbav:"SK"`
	Cell       localtype.CELL `dynamodbav:"Cell"`
	Name       string         `dynamodbav:"Name"`
	Army       int            `dynamodbav:"Army"`
	Star       int            `dynamodbav:"Star"`
	Speed      float64        `dynamodbav:"Speed"`
	Attack     float64        `dynamodbav:"Attack"`
	Face       string         `dynamodbav:"Face"`
	CreatedAt  int64          `dynamodbav:"CreatedAt"`
	EntityType string         `dynamodbav:"EntityType"`
}

func (m *Marshal) ToRepo() MarshalRepo {
	return MarshalRepo{
		UserID:     hashtag.User + m.UserID,
		ID:         hashtag.Marshal + m.ID,
		Cell:       m.Cell,
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
		UserID:    m.UserID,
		ID:        m.ID,
		Cell:      m.Cell,
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
