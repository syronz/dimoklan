package model

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"dimoklan/consts/gp"
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

type MarshalMove struct {
	MarshalID   string              `json:"marshal_id" redis:"-"`
	UserID      string              `json:"user_id" redis:"-"`
	Name        string              `json:"name" redis:"-"`
	Star        int                 `json:"star" redis:"-"`
	Speed       float64             `json:"speed" redis:"-"`
	Face        string              `json:"face" redis:"-"`
	Directrion  localtype.DIRECTION `json:"direction" redis:"-"`
	Source      localtype.CELL      `json:"source" redis:"source"`
	Destination localtype.CELL      `json:"destination" redis:"destination"`
	DepartureAt int64               `json:"departure_at" redis:"departure_at"`
	ArriveAt    int64               `json:"arrived_at" redis:"arrive_at"`
}

/*
Input:

	{
	  "marshal_id": "m:3224053:1", -- ignored
	  "user_id": "u:3224053",
	  "name": "Napoleon",
	  "star": 1,
	  "speed": 1,
	  "face": "no-face",
	  "direction": "d",
	  "source": "c:2:6",
	  "destination": "c:2:16",
	  "departure_at": 1719048815969,
	  "arrived_at": 1257897000000
	}

Output:
"u:3224053","Napoleon",1,1,"no-face","d","c:2:6","c:2:16",1719048815969,1719048815969
*/
func (mm *MarshalMove) ToZipString() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v",
		mm.UserID,
		mm.Name,
		mm.Star,
		mm.Speed,
		mm.Face,
		mm.Directrion,
		mm.Source,
		mm.Destination,
		mm.DepartureAt,
		mm.ArriveAt)
}

func ZipStringToMarshalMove(marshalID, str string) (MarshalMove, error) {
	arr := strings.Split(str, ",")

	if len(arr) != 10 {
		return MarshalMove{}, fmt.Errorf("number of values inside the zip-string is not valid; marshal_id:%v, data: %v", marshalID, str)
	}

	star, err := strconv.Atoi(arr[2])
	if err != nil {
		return MarshalMove{}, fmt.Errorf("error in converting star to number; marshalID: %v; data: %v; err:%w", marshalID, str, err)
	}

	speed, err := strconv.ParseFloat(arr[3], 64)
	if err != nil {
		return MarshalMove{}, fmt.Errorf("error in converting speed to number; marshalID: %v; data: %v; err:%w", marshalID, str, err)
	}

	direction, err := localtype.SetDirection(arr[7])
	if err != nil {
		return MarshalMove{}, fmt.Errorf("error in set direction; marshalID: %v; data: %v; err: %w", marshalID, str, err)
	}

	departureAt, err := strconv.ParseInt(arr[8], 10, 64)
	if err != nil {
		return MarshalMove{}, fmt.Errorf("error in converting departure_at to number; marshalID: %v; data: %v; err:%w", marshalID, str, err)
	}

	arrivedAt, err := strconv.ParseInt(arr[9], 10, 64)
	if err != nil {
		return MarshalMove{}, fmt.Errorf("error in converting arrived_at to number; marshalID: %v; data: %v; err:%w", marshalID, str, err)
	}

	moveMarshal := MarshalMove{
		MarshalID:   marshalID,
		UserID:      arr[0],
		Name:        arr[1],
		Star:        star,
		Speed:       speed,
		Face:        arr[4],
		Directrion:  direction,
		Source:      localtype.CELL(arr[6]),
		Destination: localtype.CELL(arr[7]),
		DepartureAt: departureAt,
		ArriveAt:    arrivedAt,
	}

	return moveMarshal, nil
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
