package service

var (
	/*
			TABLE users (
		    id SERIAL PRIMARY KEY,
		    profile JSONB
			)
	*/
	CreateUserQuery = "INSERT INTO users_table(firstname,lastname,username,email,password," +
		"passwordconfirmed,age,birthdate,phonenumber,address,education,interests,headline,intent,subscriptions,bio)" +
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)"
	CheckIfUserAlreadyExistQuery = "SELECT id FROM users_table WHERE username =$1 AND email =$2"
	GetUserByIdQuery = "SELECT * FROM users_table WHERE id =$1"
	GetUserByUsernameQuery = "SELECT * FROM users_table WHERE username =$1"
	GetUserByEmailQuery = "SELECT * FROM users_table WHERE email =$1"

)