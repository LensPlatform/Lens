package service

var (
	/*
			TABLE users (
		    id SERIAL PRIMARY KEY,
		    profile JSONB
			)
	*/
	CreateUserQuery = "INSERT INTO users_table(firstname,lastname,username,email,password," +
		"passwordconfirmed,age,birthdate,phonenumber) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	GetUserByIdQuery = "SELECT * FROM users WHERE profile->'id' == VALUES($1)"
	GetUserByUserNameQuery = "SELECT * FROM users WHERE profile->'user_name' == VALUES($1)"
	GetUserByEmailQuery = "SELECT * FROM users WHERE profile->'email' == VALUES($1)"
	CheckIfEmailAlreadyExists = "SELECT * FROM users WHERE profile->'email' == VALUES($1)"
)