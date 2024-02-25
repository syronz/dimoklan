package types

type User struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name" faker:"name"`
	Email     string `json:"email" faker:"email"`
	Kingdom   string `json:"kingdom"`
	Password  string `json:"password"`
	Color     string `json:"color"`
	Language  string `json:"language"`
	Status    string `json:"status"`
	Reason    string `json:"reason"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func validateUser(u *User) bool { return true }
