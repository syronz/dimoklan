package types

import "time"

type User struct {
	Code      string    `json:"code"`
	Bit       int       `json:"bit"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Color     string    `json:"color"`
	Language  string    `json:"language"`
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func validateUser(u *User) bool { return true }
