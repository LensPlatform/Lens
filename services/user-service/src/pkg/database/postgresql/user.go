package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/LensPlatform/Lens/services/user-service/src/pkg/helper"
	table "github.com/LensPlatform/Lens/services/user-service/src/pkg/models/proto"
)

func (db *Database) CreateUser(user table.UserORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		var foundUser table.UserORM

		ctx := context.TODO()
		pbUser, err := user.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbUser.Validate()
		if err != nil {
			return err
		}

		// check if user exists based on email or username fields
		// Note: email and username fields are unique so if a db entity witholds those respective parameters
		// we know the user already exists
		if err = tx.Where("email = ?", pbUser.Email).Or("username = ?", pbUser.UserName).Find(&foundUser).Error; err != nil {
			return err
		}

		if foundUser.Email != "" {
			return helper.ErrAlreadyExists
		}

		// save the user to the database
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) UpdateUser(user table.UserORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		ctx := context.TODO()
		pbUser, err := user.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbUser.Validate()
		if err != nil {
			return err
		}

		// Updates all fields in a user entity in the database
		if err = tx.Save(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) DeleteUser(user table.UserORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		var foundUser table.UserORM
		ctx := context.TODO()
		pbUser, err := user.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbUser.Validate()
		if err != nil {
			return err
		}

		// Check if user exists
		if err := tx.First(&foundUser, user.Id).Error; err != nil {
			return err
		}

		// user exists in the database hence perform deletion
		if err = tx.Where("email = ?", user.Email).Delete(&foundUser).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) GetUserById(id int32) (error, *table.UserORM) {
	var foundUser table.UserORM

	// attempt to obtain a user from the database with this id
	if err := db.Engine.Where("id = ?", id).First(&foundUser).Error; err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	// Validate all required fields are populated
	ctx := context.TODO()
	if err := db.validateUser(foundUser, ctx); err != nil {
		return err, nil
	}

	return nil, &foundUser
}

func (db *Database) GetUserByUsername(username string) (error, *table.UserORM) {
	var foundUser table.UserORM

	// attempt to obtain a user from the database with this username
	if err := db.Engine.Where("username = ?", username).First(&foundUser).Error; err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	// Validate all required fields are populated
	ctx := context.TODO()
	if err := db.validateUser(foundUser, ctx); err != nil {
		return err, nil
	}

	return nil, &foundUser
}

func (db *Database) GetUserByEmail(email string) (error, *table.UserORM) {
	var foundUser table.UserORM

	// attempt to obtain a user from the database with this username
	if err := db.Engine.Where("email = ?", email).First(&foundUser).Error; err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	// Validate all required fields are populated
	ctx := context.TODO()
	if err := db.validateUser(foundUser, ctx); err != nil {
		return err, nil
	}

	return nil, &foundUser
}

func (db *Database) GetAllUsers(limit int) (error, []*table.UserORM) {
	var users []*table.UserORM
	var user table.UserORM

	// find all users in the users tables but do not breach limit
	rows, err := db.Engine.Limit(limit).Find(table.UserORM{}).Rows()
	if err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	for rows.Next() {
		if err := rows.Scan(&user); err != nil {
			return err, nil
		}

		users = append(users, &user)
	}

	return nil, users
}

func (db Database) GetUserBasedOnParam(param string, query string) (table.UserORM, error) {
	if param == "" || query == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of "+
			"the following params are null Search Param : %s, Query : %s", param, query)
		return table.UserORM{}, errors.New(errMsg)
	}

	var user table.UserORM
	rows, e := db.Engine.Table(UsersTableName).Raw(query, param).Rows()
	if e != nil {
		return user, e
	}

	defer rows.Close()
	for rows.Next() {
		_ = db.Engine.ScanRows(rows, &user)
	}

	return user, e
}
