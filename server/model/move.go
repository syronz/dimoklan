package model

import (
	"fmt"
	"slices"
	"strings"

	"dimoklan/consts/gp"
	"dimoklan/consts/hashtag"
	"dimoklan/internal/errors/errstatus"
	"dimoklan/model/localtype"
)

type Move struct {
	Type      string         `json:"type"`
	Cell      localtype.CELL `json:"cell"`
	MarshalID string         `json:"marshal_id"`
	UserID    string         `json:"-"`
}

func (m *Move) GetPkSkforMoving() (pk string,sk string) {
	return hashtag.Fraction + m.Cell.ToFraction(), hashtag.MarshalEx + m.MarshalID
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
	if len(parsedMarshalID) != 2 {
		return fmt.Errorf("marshal_id is not valid; code: %w", errstatus.ErrNotAcceptable)
	}

	c.UserID = parsedMarshalID[0]

	if err := c.Cell.Validate(); err != nil {
		return err
	}

	return nil
}
