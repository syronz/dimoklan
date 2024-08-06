package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go.uber.org/zap"

	"dimoklan/internal/config"
	"dimoklan/model/localtype"
)

type Fraction struct {
	Fraction   string         `json:"fraction,omitempty" dynamodbav:"PK"`
	CellStr    string         `json:"cell,omitempty" dynamodbav:"SK"`
	Cell       localtype.CELL `json:"cell,omitempty" dynamodbav:"-"`
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

func (m *Fraction) GetKey(core config.Core) map[string]types.AttributeValue {
	pk, err := attributevalue.Marshal(m.Fraction)
	if err != nil {
		core.Error("DANGER: failed to marshal fraction pk", zap.Error(err), zap.Stack("marshal_fraction_pk_failed"))
	}
	sk, err := attributevalue.Marshal(m.Cell)
	if err != nil {
		core.Error("DANGER: failed to marshal fraction sk", zap.Error(err), zap.Stack("marshal_fraction_sk"))
	}
	return map[string]types.AttributeValue{"PK": pk, "SK": sk}
}
