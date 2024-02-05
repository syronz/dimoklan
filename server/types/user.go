package types

import "time"

type User struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name" faker:"name"`
	Email     string    `json:"email" faker:"email"`
	Username  string    `json:"username" faker:"username"`
	Password  string    `json:"password"`
	Color     string    `json:"color"`
	Language  string    `json:"language"`
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func validateUser(u *User) bool { return true }
