package types

type User struct {
	ID            int    `json:"id"`
	Color         string `json:"color"`
	Email         string `json:"email"`
	Kingdom       string `json:"kingdom"`
	Password      string `json:"password"`
	Language      string `json:"language"`
	Suspend       bool   `json:"suspend"`
	SuspendReason string `json:"suspend_reason"`
	Freeze        bool   `json:"freeze"`
	FreezeReason  string `json:"freeze_reason"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

func validateUser(u *User) bool { return true }
