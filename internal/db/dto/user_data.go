package dto

type UserData struct {
	ID         string `db:"id"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Age        int    `db:"age"`
	Sex        string `db:"sex"`
	Biography  string `db:"biography"`
	City       string `db:"city"`
}
