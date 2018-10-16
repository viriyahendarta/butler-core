package user

const queryGetUserByID = `
	SELECT * FROM users
	WHERE user_id = $1
	LIMIT 1
`
