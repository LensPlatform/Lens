package postgresql

import (
	"context"
	"errors"

	table "github.com/LensPlatform/Lens/services/user-service/src/pkg/models/proto"
)

func (db *Database) validateUser(user table.UserORM, ctx context.Context) error {

	if user.Email == "" {
		return errors.New("User Does Not Exist")
	}
	pbUser, err := user.ToPB(ctx)

	if err != nil {
		db.Logger.Error(err.Error())
		return err
	}

	if err := pbUser.Validate(); err != nil {
		db.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (db *Database) validateTeam(team table.TeamORM, ctx context.Context) error {

	if team.TeamEmail == "" {
		return errors.New("team does not exist")
	}
	pbTeam, err := team.ToPB(ctx)

	if err != nil {
		db.Logger.Error(err.Error())
		return err
	}

	if err := pbTeam.Validate(); err != nil {
		db.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (db *Database) validateGroup(group table.GroupORM, ctx context.Context) error {

	if group.GroupName == "" {
		return errors.New("group does not exist")
	}
	pbGroup, err := group.ToPB(ctx)

	if err != nil {
		db.Logger.Error(err.Error())
		return err
	}

	if err := pbGroup.Validate(); err != nil {
		db.Logger.Error(err.Error())
		return err
	}
	return nil
}
