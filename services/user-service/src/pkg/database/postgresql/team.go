package postgresql

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"

	table "github.com/LensPlatform/Lens/services/user-service/src/pkg/models/proto"
)

func (db *Database) CreateTeam(team table.TeamORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		var foundTeam table.TeamORM

		ctx := context.TODO()
		pbTeam, err := team.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbTeam.Validate()
		if err != nil {
			return err
		}

		// check if team exists based on email or teamname fields
		// Note: email and username fields are unique so if a db entity witholds those respective parameters
		// we know the team already exists
		if err = tx.Where("email = ?", pbTeam.Email).Or("teamName = ?", pbTeam.Name).Find(&foundTeam).Error; err != nil {
			return err
		}

		if foundTeam.Email != "" {
			return errors.New("team already exists")
		}

		// save the user to the database
		if err := tx.Create(team).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) UpdateTeam(team table.TeamORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		ctx := context.TODO()
		pbTeam, err := team.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbTeam.Validate()
		if err != nil {
			return err
		}

		// Updates all fields in a team entity in the database
		if err = tx.Save(team).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) DeleteTeam(team table.TeamORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		var foundTeam table.TeamORM
		ctx := context.TODO()
		pbUser, err := team.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbUser.Validate()
		if err != nil {
			return err
		}

		// Check if team exists
		if err := tx.First(&foundTeam, team.Id).Error; err != nil {
			return err
		}

		// user exists in the database hence perform deletion
		if err = tx.Where("email = ?", team.Email).Delete(&foundTeam).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) GetTeamById(id int32) (error, *table.TeamORM) {
	var foundTeam table.TeamORM

	// attempt to obtain a user from the database with this id
	if err := db.Engine.Where("id = ?", id).First(&foundTeam).Error; err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	// Validate all required fields are populated
	ctx := context.TODO()
	if err := db.validateTeam(foundTeam, ctx); err != nil {
		return err, nil
	}

	return nil, &foundTeam
}

func (db *Database) GetTeamByName(teamName string) (error, *table.TeamORM) {
	var foundTeam table.TeamORM

	// attempt to obtain a user from the database with this username
	if err := db.Engine.Where("teamname = ?", teamName).First(&foundTeam).Error; err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	// Validate all required fields are populated
	ctx := context.TODO()
	if err := db.validateTeam(foundTeam, ctx); err != nil {
		return err, nil
	}

	return nil, &foundTeam
}

func (db *Database) GetAllTeams(limit int) (error, []*table.TeamORM) {
	var teams []*table.TeamORM
	var team table.TeamORM

	// find all teams in the teams tables but do not breach limit
	rows, err := db.Engine.Limit(limit).Find(table.TeamORM{}).Rows()
	if err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	for rows.Next() {
		if err := rows.Scan(&team); err != nil {
			return err, nil
		}

		teams = append(teams, &team)
	}

	return nil, teams
}
