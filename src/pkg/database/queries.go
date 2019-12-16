package database

var (
	GetUserByIdQuery       = "SELECT * FROM users_table WHERE id =$1"
	GetUserByUsernameQuery = "SELECT * FROM users_table WHERE username =$1"
	GetUserByEmailQuery    = "SELECT * FROM users_table WHERE email =$1"
)
