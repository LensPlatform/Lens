package models

// User represents a single user profile
// ID should always be globally unique
type User struct {
	ID string `json:"id" validate:"-" sql:"id"`
	FirstName string `json:"first_name" validate:"required" sql:"firstname"`
	LastName string `json:"last_name" validate:"required" sql:"lastname"`
	UserName string `json:"user_name" validate:"required" sql:"username"`
	Gender string `json:"gender" validate:"-" sql:"gender"`
	Languages string `json:"Languages" validate:"-" sql:"languages"`
	Email string `json:"email" validate:"required,email" sql:"email"`
	PassWord string `json:"password" validate:"required,gte=8,lte=20" sql:"password"`
	PassWordConfirmed string `json:"password_confirmed" validate:"required,gte=8,lte=20" sql:"passwordconf"`
	Age int `json:"age" validate:"gte=0,lte=120" sql:"age"`
	BirthDate string `json:"birth_date" validate:"required" sql:"birthdate"`
	PhoneNumber string `json:"phone_number" validate:"required" sql:"phonenumber"`
	Addresses Address `json:"location" validate:"-" sql:"address"`
	Bio string `json:"bio" validate:"required" sql:"bio"`
	Education Education `json:"education" validate:"-" sql:"education"`
	UserInterests Interests `json:"interests" validate:"-" sql:"interests"`
	Headline string `json:"headline" validate:"max=30" sql:"headline"`
	Subscriptions Subscriptions `json:"subscriptions" validate:"-" sql:"subscriptions"`
	Intent string `json:"intent" validate:"required" sql:"intent"`
	Skills Skillset `json:"skillset" validate:"-" sql:"skills"`
	Teams []string `json:"associated_teams" validate:"-" sql:"teams"`
	Groups []string `json: "associated_groups" validate:"-" sql: "groups"`
	SocialMedia SocialMedia `json:"social_media" validate:"-" sql:"social_media"`
	Settings Settings `json:"settings" validate:"-" sql:"settings"`
}

type Address struct {
	JsonEmbeddable
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




