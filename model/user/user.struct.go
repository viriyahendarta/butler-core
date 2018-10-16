package user

type User struct {
	ID        string `db:"user_id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}
