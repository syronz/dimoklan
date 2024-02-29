package types

type User struct {
	ID            string `json:"id" dynamodbav:"PK"`
	Color         string `json:"color" dynamodbav:"Color"`
	Farr          int    `json:"farr" dynamodbav:"Farr"`
	Gold          int    `json:"gold" dynamodbav:"Gold"`
	Email         string `json:"email" dynamodbav:"Email"`
	Kingdom       string `json:"kingdom" dynamodbav:"Kingdom"`
	Password      string `json:"password,omitempty" dynamodbav:"Password"`
	Language      string `json:"language" dynamodbav:"Language"`
	Suspend       bool   `json:"suspend" dynamodbav:"Suspend"`
	SuspendReason string `json:"suspend_reason" dynamodbav:"SuspendReason"`
	Freeze        bool   `json:"freeze" dynamodbav:"Freeze"`
	FreezeReason  string `json:"freeze_reason" dynamodbav:"FreezeReason"`
	CreatedAt     int64  `json:"created_at" dynamodbav:"CreatedAt"`
	UpdatedAt     int64  `json:"updated_at" dynamodbav:"UpdatedAt"`
	SK            string `json:"-" dynamodbav:"SK"`
	EntityType    string `json:"-" dynamodbav:"EntityType"`
}

func validateUser(u *User) bool { return true }
