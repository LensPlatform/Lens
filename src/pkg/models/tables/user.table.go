package tables

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
)

type UserTable struct{
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Type string `json:"user_type" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	UserName string `json:"user_name" validate:"required" gorm:"type:varchar(100);unique_index"`
	Gender string `json:"gender" validate:"-"`
	Languages string `json:"Languages" validate:"-"`
	Email string `json:"email" validate:"required,email"`
	PassWord string `json:"password" validate:"required,gte=8,lte=20"`
	PassWordConfirmed string `json:"password_confirmed" validate:"required,gte=8,lte=20"`
	Age int `json:"age" validate:"gte=0,lte=120"`
	BirthDate string `json:"birth_date" validate:"required"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required"`
	Addresses postgres.Jsonb `json:"location,omitempty" validate:"-"`
	Bio string `json:"bio,omitempty" validate:"required"`
	Education postgres.Jsonb `json:"education,omitempty" validate:"-"`
	UserInterests postgres.Jsonb `json:"interests,omitempty" validate:"-"`
	Headline string `json:"headline,omitempty" validate:"max=30"`
	Subscriptions postgres.Jsonb `json:"subscriptions,omitempty" validate:"-"`
	Intent string `json:"intent,omitempty" validate:"required"`
	Skills postgres.Jsonb `json:"skillset,omitempty" validate:"-"`
	Teams postgres.Jsonb `json:"associated_teams,omitempty" validate:"-"`
	Groups postgres.Jsonb `json: "associated_groups,omitempty" validate:"-"`
	SocialMedia postgres.Jsonb `json:"social_media,omitempty" validate:"-"`
	Settings postgres.Jsonb `json:"settings,omitempty" validate:"-"`
}

func (table UserTable) MigrateSchemaOrCreateTable(db *gorm.DB, logger *zap.Logger){
	t := reflect.TypeOf(table)
	tableName := t.Name()

	if db.HasTable(&table) {
		err := db.AutoMigrate(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Migrate: %s Schema", tableName) )
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Successfully Migrated %s Schema", tableName) )
	} else {
		err := db.CreateTable(&table).Error
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot Create %s Table", tableName))
			logger.Error(err.Error())
		}
		logger.Info(fmt.Sprintf("Sucessfully Created %s Table", tableName))
	}
}