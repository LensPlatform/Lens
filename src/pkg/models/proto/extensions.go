package user

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func (table UserORM) MigrateSchemaOrCreateTable(db *gorm.DB, logger *zap.Logger) {
	t := reflect.TypeOf(table)
	tableName := t.Name()

	if db.HasTable(&table) {
		err := db.AutoMigrate(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Migrate: %s Schema", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Successfully Migrated %s Schema", tableName))
	} else {
		err := db.CreateTable(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Create %s Table", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Sucessfully Created %s Table", tableName))
	}
}

func (table TeamORM) MigrateSchemaOrCreateTable(db *gorm.DB, logger *zap.Logger) {
	t := reflect.TypeOf(table)
	tableName := t.Name()

	if db.HasTable(&table) {
		err := db.AutoMigrate(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Migrate: %s Schema", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Successfully Migrated %s Schema", tableName))
	} else {
		err := db.CreateTable(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Create %s Table", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Sucessfully Created %s Table", tableName))
	}
}

func (table GroupORM) MigrateSchemaOrCreateTable(db *gorm.DB, logger *zap.Logger) {
	t := reflect.TypeOf(table)
	tableName := t.Name()

	if db.HasTable(&table) {
		err := db.AutoMigrate(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Migrate: %s Schema", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Successfully Migrated %s Schema", tableName))
	} else {
		err := db.CreateTable(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Create %s Table", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Sucessfully Created %s Table", tableName))
	}
}

func (table AddressORM) MigrateSchemaOrCreateTable(db *gorm.DB, logger *zap.Logger) {
	t := reflect.TypeOf(table)
	tableName := t.Name()

	if db.HasTable(&table) {
		err := db.AutoMigrate(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Migrate: %s Schema", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Successfully Migrated %s Schema", tableName))
	} else {
		err := db.CreateTable(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Create %s Table", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Sucessfully Created %s Table", tableName))
	}
}


