package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBHandler interface {
	AddUser(user User) error
	GetUserById(id string) ( User,  error)
	GetUserByEmail(email string)( User,  error)
	GetUserByUsername(username string)( User,  error)
	GetPassword(id string)( string,  error)
	DoesUserExist(id string) (bool, error)
}

type Database struct {
	connection *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{connection:db}
}

func (db Database) AddUser(user User) error {
	// Create user in database
	_, err := db.connection.Exec(CreateUserQuery, user.FirstName, user.LastName,
		user.UserName, user.Email, user.PassWord,
		user.PassWordConfirmed,user.Age, user.BirthDate,
		user.PhoneNumber, user.Addresses, user.EducationalExperience,
		user.UserInterests,user.Headline,user.Intent,
		user.UserSubscriptions, user.Bio, user.Gender,
		user.Skills, user.Languages)

	if err != nil {
		return err
	}

	return nil
}

func (db Database) GetUserById(id string) (User,error) {
	var user User

	// Obtain user by id
	user, err := db.GetUserBasedOnParam(id,GetUserByIdQuery)

	if err != nil {
		return User{},err
	}

	return user, nil
}

func (db Database) GetUserByUsername(username string) (User,error) {
	var user User

	// Obtain user by id
	user, err := db.GetUserBasedOnParam(username,GetUserByUsernameQuery)

	if err != nil {
		return User{},err
	}

	return user, nil
}

func (db Database) GetUserByEmail(email string) (User,error) {
	var user User

	// Obtain user by id
	user, err := db.GetUserBasedOnParam(email,GetUserByEmailQuery)

	if err != nil {
		return User{},err
	}

	return user, nil
}

func (db Database) GetUserBasedOnParam(param string, query string) (User, error) {
		var user User
		rows := db.connection.QueryRow(query, param)

		if rows == nil{
			return User{}, nil
		}

		err := rows.Scan(&user.ID,&user.FirstName, &user.LastName, &user.UserName, &user.Email,
								&user.PassWord, &user.PassWordConfirmed,
								&user.Age, &user.BirthDate, &user.PhoneNumber,&user.Addresses,
								&user.Bio, &user.EducationalExperience,
								&user.UserInterests, &user.Headline, &user.Intent, &user.UserSubscriptions,
								&user.Gender, &user.Languages,&user.Skills)

		if err != nil{

			return User{},err
		}

		return user, nil
}

func (db Database) DoesUserExist(username string) (bool,error) {
	// check if user exists
	var user User
	rows := db.connection.QueryRow(GetUserByUsernameQuery, username)

	err := rows.Scan(&user.ID,&user.FirstName, &user.LastName, &user.UserName, &user.Email,
		&user.PassWord, &user.PassWordConfirmed,
		&user.Age, &user.BirthDate, &user.PhoneNumber,&user.Addresses,
		&user.Bio, &user.EducationalExperience,
		&user.UserInterests, &user.Headline, &user.Intent, &user.UserSubscriptions,
		&user.Gender, &user.Languages,&user.Skills)

	if err != nil {
		if err == sql.ErrNoRows{
			return false, nil
		}
		return true, err
	}

	if user.ID != ""{
		return true, nil
	}

	return false, nil
}

