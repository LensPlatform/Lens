package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User represents a single user profile
// ID should always be globally unique
type User struct {
	JsonEmbeddable
	ID string `json:"id" validate:"-" gorm:"primary_key"`
	Type string `json:"user_type" validate:"required"`
	FirstName string `json:"first_name" validate:"required" sql:"firstname"`
	LastName string `json:"last_name" validate:"required" sql:"lastname"`
	UserName string `json:"user_name" validate:"required" gorm:"type:varchar(100);unique_index"`
	Gender string `json:"gender" validate:"-" sql:"gender"`
	Languages string `json:"Languages" validate:"-" sql:"languages"`
	Email string `json:"email" validate:"required,email" gorm:"type:varchar(100);unique_index"`
	PassWord string `json:"password" validate:"required,gte=8,lte=20" sql:"password"`
	PassWordConfirmed string `json:"password_confirmed" validate:"required,gte=8,lte=20" sql:"passwordconf"`
	Age int `json:"age" validate:"gte=0,lte=120" sql:"age"`
	BirthDate string `json:"birth_date" validate:"required" sql:"birthdate"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required" sql:"phonenumber"`
	Addresses Address `json:"location,omitempty" validate:"-" sql:"address"`
	Bio string `json:"bio,omitempty" validate:"required" sql:"bio"`
	Education Education `json:"education,omitempty" validate:"-" sql:"education"`
	UserInterests Interests `json:"interests,omitempty" validate:"-" sql:"interests"`
	Headline string `json:"headline,omitempty" validate:"max=30" sql:"headline"`
	Subscriptions Subscriptions `json:"subscriptions,omitempty" validate:"-" sql:"subscriptions"`
	Intent string `json:"intent,omitempty" validate:"required" sql:"intent"`
	Skills Skillset `json:"skillset,omitempty" validate:"-" sql:"skills"`
	Teams []string `json:"associated_teams,omitempty" validate:"-" sql:"teams"`
	Groups []string `json: "associated_groups,omitempty" validate:"-" sql: "groups"`
	SocialMedia SocialMedia `json:"social_media,omitempty" validate:"-" sql:"social_media"`
	Settings Settings `json:"settings,omitempty" validate:"-" sql:"settings"`
	CreatedAt string `json:"created_at" validate:"-"`
	UpdatedAt string `json:"updated_at" validate:"-"`
}

type Address struct {
	JsonEmbeddable
	gorm.Model
	City string `json:"city" validate:"required" sql:"city"`
	State string `json:"state" validate:"required" sql:"state"`
	Country string `json:"country" validate:"required" sql:"country"`
}

type Education struct{
	JsonEmbeddable
	gorm.Model
	MostRecentInstitutionName string `json:"most_recent_institution_name" validate:"required"`
	HighestDegreeEarned string `json:"highest_degree_earned" validate:"required"`
	Graduated bool `json:"graduated" validate:"required"`
	Major string `json:"major" validate:"required"`
	Minor string `json:"minor" validate:"required"`
	YearsOfAttendance string `json:"years_of_attendance" validate:"required"`
}

type Interests struct {
	JsonEmbeddable
	gorm.Model
	Industry []Industry `json:"industries_of_interest" validate:"omitempty"`
	Topic []Topic `json:"topics_of_interest" validate:"omitempty"`
}

type Topic struct{
	JsonEmbeddable
	gorm.Model
	TopicName string `json:"topic_name" validate:"required"`
	TopicType string `json:"topic_type" validate:"required"`
}

type Industry struct {
	JsonEmbeddable
	gorm.Model
	IndustryName string `json:"industry_name" validate:"required"`
}

type Subscriptions struct {
	JsonEmbeddable
	gorm.Model
	SubscriptionName string `json:"subscription_name" validate:"required"`
	Subscribe bool `json:"subscribe" validate:"required"`
}

type Skillset struct {
	JsonEmbeddable
	gorm.Model
	Skills []Skill `json:"skills" validate:"required"`
}

type Skill struct {
	JsonEmbeddable
	gorm.Model
	Type string `json:"skill_type" validate:"required"`
	Name string `json:"skill_name" validate:"required"`
}

func (user User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New()
	err := scope.SetColumn("createdat", time.Now().String())

	if err != nil{
		return err
	}

	err = scope.SetColumn("updatedat", time.Now().String())

	if err != nil{
		return err
	}

	return scope.SetColumn("id", id.String())
}

func (user User) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("updatedat", time.Now().String())

	if err != nil{
		return err
	}
	return nil
}