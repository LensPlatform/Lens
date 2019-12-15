package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User represents a single user profile
// ID should always be globally unique
type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CurrentID int `json:"user_id" validate:"-"`
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
	Addresses Address `json:"location,omitempty" validate:"-"`
	Bio string `json:"bio,omitempty" validate:"required"`
	Education Education `json:"education,omitempty" validate:"-"`
	UserInterests Interests `json:"interests,omitempty" validate:"-"`
	Headline string `json:"headline,omitempty" validate:"max=30"`
	Subscriptions Subscriptions `json:"subscriptions,omitempty" validate:"-"`
	Intent string `json:"intent,omitempty" validate:"required"`
	Skills Skillset `json:"skillset,omitempty" validate:"-"`
	Teams []string `json:"associated_teams,omitempty" validate:"-"`
	Groups []string `json: "associated_groups,omitempty" validate:"-"`
	SocialMedia SocialMedia `json:"social_media,omitempty" validate:"-"`
	Settings Settings `json:"settings,omitempty" validate:"-"`
}

type Address struct {
	City string `json:"city" validate:"required" sql:"city"`
	State string `json:"state" validate:"required" sql:"state"`
	Country string `json:"country" validate:"required" sql:"country"`
}

type Education struct{
	JsonEmbeddable
	MostRecentInstitutionName string `json:"most_recent_institution_name" validate:"required"`
	HighestDegreeEarned string `json:"highest_degree_earned" validate:"required"`
	Graduated bool `json:"graduated" validate:"required"`
	Major string `json:"major" validate:"required"`
	Minor string `json:"minor" validate:"required"`
	YearsOfAttendance string `json:"years_of_attendance" validate:"required"`
}

type Interests struct {
	JsonEmbeddable
	Industry []Industry `json:"industries_of_interest" validate:"omitempty"`
	Topic []Topic `json:"topics_of_interest" validate:"omitempty"`
}

type Topic struct{
	JsonEmbeddable
	TopicName string `json:"topic_name" validate:"required"`
	TopicType string `json:"topic_type" validate:"required"`
}

type Industry struct {
	JsonEmbeddable
	IndustryName string `json:"industry_name" validate:"required"`
}

type Subscriptions struct {
	JsonEmbeddable
	SubscriptionName string `json:"subscription_name" validate:"required"`
	Subscribe bool `json:"subscribe" validate:"required"`
}

type Skillset struct {
	JsonEmbeddable
	Skills []Skill `json:"skills" validate:"required"`
}

type Skill struct {
	JsonEmbeddable
	Type string `json:"skill_type" validate:"required"`
	Name string `json:"skill_name" validate:"required"`
}

func (user User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New()
	err := scope.SetColumn("created_at", time.Now())

	if err != nil{
		return err
	}

	err = scope.SetColumn("updated_at", time.Now())

	if err != nil{
		return err
	}

	return scope.SetColumn("id", id)
}

func (user User) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("updatedat", time.Now().String())

	if err != nil{
		return err
	}
	return nil
}

