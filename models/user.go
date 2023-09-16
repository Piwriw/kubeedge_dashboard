package models

type User struct {
	UserID int64 `db:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Token string
}
