package dto

type Credentials struct {
	ID       string `db:"id"`
	UserID   string `db:"user_id"`
	Password string `db:"password"`
}
