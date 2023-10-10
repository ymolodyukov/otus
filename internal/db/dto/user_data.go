package dto

type UserData struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int    `db:"age"`
	Sex       string `db:"sex"`
	Biography string `db:"biography"`
	City      string `db:"city"`
}
