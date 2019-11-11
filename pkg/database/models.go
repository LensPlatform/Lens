package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"google.golang.org/genproto/googleapis/type/date"
)

type JsonEmbeddable struct {}

// User represents a single user profile
// ID should always be globally unique
type User struct {
	ID string `json:"id" validate:"-" db:"id"`
	FirstName string `json:"first_name" validate:"required" db:"firstname"`
	LastName string `json:"last_name" validate:"required" db:"lastname"`
	UserName string `json:"user_name" validate:"required" db:"username"`
	Gender string `json:"gender" validate:"-" db:"gender"`
	Languages string `json:"Languages" validate:"-" db:"languages"`
	Email string `json:"email" validate:"required,email" db:"email"`
	PassWord string `json:"password" validate:"required,gte=8,lte=20" db:"password"`
	PassWordConfirmed string `json:"password_confirmed" validate:"required,gte=8,lte=20" db:"passwordconf"`
	Age int `json:"age" validate:"gte=0,lte=120" db:"age"`
	BirthDate string `json:"birth_date" validate:"required" db:"birthdate"`
	PhoneNumber string `json:"phone_number" validate:"required" db:"phonenumber"`
	Addresses Address `json:"location" validate:"-" db:"address"`
	Bio string `json:"bio" validate:"required" db:"bio"`
	EducationalExperience Education `json:"education" validate:"-" db:"education"`
	UserInterests Interests `json:"interests" validate:"-" db:"interests"`
	Headline string `json:"headline" validate:"max=30" db:"headline"`
	UserSubscriptions Subscriptions `json:"subscriptions" validate:"-" db:"subscriptions"`
	Intent string `json:"intent" validate:"required" db:"intent"`
	Skills Skillset `json:"skillset" validate:"-" db:"skills"`
	AssociatedTeams []string `json:"associated_teams" validate:"-" db:"associated_teams"`
}

type Team struct{
	Name string `json:"name" validate:"required"` // team name
	Type string `json:"type" validate:"required"` // investor or startup team
	Overview string `json:"overview" validate:"required"` // about the team
	IndustryOfInterest string `json:"industry" validate:"required"` // industry of interest
	FoundedDate date.Date `json:"founded_date" validate:"required"`
	Founders []TeamMember `json:"founder" validate:"required"`
	NumberOfEmployees int `json:"number_of_employees" validate:"required"` // size of team
	Headquarters string `json:"headquarters" validate:"-"`
	Interests string `json:"interests" validate:"-"`
	TeamMembers []TeamMember `json:"team_members" validate:"-"`
	Advisors []TeamMember `json:"advisors" validate:"-"`
	SocialMedia SocialMedia `json:"social_media" validate:"-"`
	Contact Contact `json:"contact" validate:"-"`
}

type Startup struct {
	Team
	Funds Funding `json:"funding" validate:"-"`
	CompanyDetails Details `json:"company_details" validate:"-"`
}

type Investor struct{
	Team
	InvestorType string `json:"investor_type" validate:"-"`
	InvestmentStage string `json:"investment_stage" validate:"-"`
	NumberOfExits int `json:"number_of_exits" validate:"-"`
	NumberOfinvestments int `json:"number_of_investments" validate:"-"`
	NumberOfFunds int `json:"number_of_funds" validate:"-"`
	TotalFundsRaised int `json:"total_funds_raised" validate:"-"`
}

type Funding struct {
	JsonEmbeddable
	TotalFunding int `json:"total_funding" validate:"required"`
	LatestRound string `json:"latest_round" validate:"required"`
	LatestRoundDate string `json:"latest_round_date" validate:"required"`
	LatestRoundFunding int `json:"latest_round_funding" validate:"required"`
	Investors []string `json:"investors" validate:"-"`
	InvestorType string `json:"investor_type" validate:"-"` //Accelerator
	InvestmentStage string `json:"investment_stage" validate:"-"` // Debt, Early Stage Venture, Seed
}

type Details struct {
	JsonEmbeddable
	IPOStatus string `json:"ipo_status" validate:"-"`
	CompanyType string `json:"company_type" validate:"-"` // for-profit
}

type TeamMember struct {
	JsonEmbeddable
	ID string `json:"ID" validate:"required"`
	Name string `json:"name" validate:"required"`
	Status string `json:"status" validate:"required"` // co-founder
}

type Contact struct {
	JsonEmbeddable
	Email string `json:"email" validate:"required"`
	PhoneNumber string `json:"phonenumber" validated:"required"`
}

type SocialMedia struct{
	JsonEmbeddable
	Website string `json:"website" validate:"-"`
	Facebook string `json:"facebook" validate:"-"`
	Twitter string `json:"twitter" validate:"-"`
	LinkedIn string `json:"linkedIn" validate:""`
}

type Address struct {
	JsonEmbeddable
	City string `json:"city" validate:"required"`
	State string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
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

// ser struct implements the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (u JsonEmbeddable) Value() (driver.Value, error) {
	return json.Marshal(u)
}

// User struct implements the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (u JsonEmbeddable) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &u)
}