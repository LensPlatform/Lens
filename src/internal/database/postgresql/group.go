package postgresql

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"

	table "github.com/LensPlatform/Lens/src/pkg/models/proto"
)

func (db *Database) CreateGroup(group table.GroupORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		var foundGroup table.GroupORM

		ctx := context.TODO()
		pbGroup, err := group.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbGroup.Validate()
		if err != nil {
			return err
		}

		// check if group exists based on id or group name fields
		// Note: email and username fields are unique so if a db entity witholds those respective parameters
		// we know the team already exists
		if err = tx.Where("Id = ?", foundGroup.Id).Or("GroupName = ?", pbGroup.GroupName).Find(&foundGroup).Error; err != nil {
			return err
		}

		if foundGroup.GroupName != "" {
			return errors.New("Group already exists")
		}

		// save the user to the database
		if err := tx.Create(group).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) UpdateGroup(group table.GroupORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		ctx := context.TODO()
		pbTeam, err := group.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbTeam.Validate()
		if err != nil {
			return err
		}

		// Updates all fields in a group entity in the database
		if err = tx.Save(group).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) DeleteGroup(group table.GroupORM) error {
	err := db.Engine.Transaction(func(tx *gorm.DB) error {
		var foundGroup table.GroupORM
		ctx := context.TODO()
		pbGroup, err := group.ToPB(ctx)
		if err != nil {
			return err
		}

		err = pbGroup.Validate()
		if err != nil {
			return err
		}

		// Check if group exists
		if err := tx.First(&foundGroup, group.Id).Error; err != nil {
			return err
		}

		// group exists in the database hence perform deletion
		if err = tx.Where("name = ?", group.GroupName).Delete(&foundGroup).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		db.Logger.Error(err.Error())
	}

	return nil
}

func (db *Database) GetGroupById(id int32) (error, *table.GroupORM) {
	var foundGroup table.GroupORM

	// attempt to obtain a group from the database with this id
	if err := db.Engine.Where("id = ?", id).First(&foundGroup).Error; err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	// Validate all required fields are populated
	ctx := context.TODO()
	if err := db.validateGroup(foundGroup, ctx); err != nil {
		return err, nil
	}

	return nil, &foundGroup
}

func (db *Database) GetGroupByName(groupname string) (error, *table.GroupORM) {
	var foundGroup table.GroupORM

	// attempt to obtain a group from the database with this username
	if err := db.Engine.Where("groupname = ?", groupname).First(&foundGroup).Error; err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	// Validate all required fields are populated
	ctx := context.TODO()
	if err := db.validateGroup(foundGroup, ctx); err != nil {
		return err, nil
	}

	return nil, &foundGroup
}

func (db *Database) GetAllGroups(limit int) (error, []*table.GroupORM) {
	var groups []*table.GroupORM
	var group table.GroupORM

	// find all groups in the groups tables but do not breach limit
	rows, err := db.Engine.Limit(limit).Find(table.GroupORM{}).Rows()
	if err != nil {
		db.Logger.Error(err.Error())
		return err, nil
	}

	for rows.Next() {
		if err := rows.Scan(&group); err != nil {
			return err, nil
		}

		groups = append(groups, &group)
	}

	return nil, groups
}
