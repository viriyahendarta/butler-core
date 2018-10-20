package user

type User struct {
	ID        string `db:"user_id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}
