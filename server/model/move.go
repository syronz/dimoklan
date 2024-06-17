package model

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"dimoklan/consts/gp"
	"dimoklan/consts/hashtag"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/model/localtype"
	"dimoklan/util"
)

type Move struct {
	Type      string         `json:"type"`
	Cell      localtype.CELL `json:"cell"`
	MarshalID string         `json:"marshal_id"`
	UserID    string         `json:"-"`
}

type MoveMarshal struct {
	MarshalID   string    `json:"marshal_id" redis:"-"`
	UserID      string    `json:"user_id" redis:"-"`
	Name        string    `json:"name" redis:"-"`
	Star        int       `json:"star" redis:"-"`
	Speed       float64   `json:"speed" redis:"-"`
	Face        string    `json:"face" redis:"-"`
	Directrion  string    `json:"direction" redis:"-"`
	Source      string    `json:"source" redis:"source"`
	Destination string    `json:"destination" redis:"destination"`
	DepartureAt time.Time `json:"departure_at" redis:"departure_at"`
	ArriveAt   time.Time `json:"arrived_at" redis:"arrive_at"`
}

func (m *Move) GetPkSkforMoving() (pk string, sk string) {
	return m.Cell.ToFractionID(), hashtag.MarshalEx + m.MarshalID
}

func (c *Move) Validate() error {
	if c.Type == "" {
		return fmt.Errorf("type is required; code: %w", errstatus.ErrNotAcceptable)
	}

	if c.MarshalID == "" {
		return fmt.Errorf("marshal_id is required; code: %w", errstatus.ErrNotAcceptable)
	}

	if c.Cell == "" {
		return fmt.Errorf("cell is required; code: %w", errstatus.ErrNotAcceptable)
	}

	if !slices.Contains(gp.MoveTypes(), c.Type) {
		return fmt.Errorf("type is not accepted; code: %w", errstatus.ErrNotAcceptable)
	}

	parsedMarshalID := strings.Split(c.MarshalID, ":")
	if len(parsedMarshalID) != 3 {
		return fmt.Errorf("marshal_id is not valid; code: %w", errstatus.ErrNotAcceptable)
	}

	c.UserID = util.ExtractUserIDFromMarshalID(c.MarshalID)

	if err := c.Cell.Validate(); err != nil {
		return err
	}

	return nil
}
