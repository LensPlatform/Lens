package service

var (
	/*
			TABLE users (
		    id SERIAL PRIMARY KEY,
		    profile JSONB
			)
	"select p.name from people as p where p.id = :id;", sql.Named("id", id)

	*/
	CreateUserQuery = "INSERT INTO users_table(firstname,lastname,username,email,password," +
		"passwordconfirmed,age,birthdate,phonenumber,address,education,interests,headline,intent,subscriptions,bio)" +
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)"
	CheckIfUserAlreadyExistQuery = "SELECT id FROM users_table WHERE username =$1 AND email =$2"
	GetUserByUserNameQuery = "SELECT * FROM users WHERE profile->'user_name' == VALUES($1)"
)