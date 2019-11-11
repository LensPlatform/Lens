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
	EducationalExperience Education `json:"education" validate:"-" sql:"education"`
	UserInterests Interests `json:"interests" validate:"-" sql:"interests"`
	Headline string `json:"headline" validate:"max=30" sql:"headline"`
	UserSubscriptions Subscriptions `json:"subscriptions" validate:"-" sql:"subscriptions"`
	Intent string `json:"intent" validate:"required" sql:"intent"`
	Skills Skillset `json:"skillset" validate:"-" sql:"skills"`
	AssociatedTeams []string `json:"associated_teams" validate:"-" sql:"associated_teams"`
}

type Team struct{
	JsonEmbeddable
	Name string `json:"name" validate:"required" sql:"name"` // team name
	Type string `json:"type" validate:"required" sql:"type"` // investor or startup team
	Overview string `json:"overview" validate:"required" sql:"overview"` // about the team
	IndustryOfInterest string `json:"industry" validate:"required" sql:"industryofinterest"` // industry of interest
	FoundedDate date.Date `json:"founded_date" validate:"required" sql:"foundeddate"`
	Founders []TeamMember `json:"founder" validate:"required" sql:"founders"`
	NumberOfEmployees int `json:"number_of_employees" validate:"required" sql:"numberofemployees"` // size of team
	Headquarters string `json:"headquarters" validate:"-" sql:"headquarters"`
	Interests string `json:"interests" validate:"-" sql:"interests"`
	TeamMembers []TeamMember `json:"team_members" validate:"-" sql:"teammembers"`
	Advisors []TeamMember `json:"advisors" validate:"-" sql:"advisors"`
	SocialMedia SocialMedia `json:"social_media" validate:"-" sql:"socialmedia"`
	Contact Contact `json:"contact" validate:"-" sql:"contact"`
}

type Startup struct {
	Team
	Funding Funding `json:"funding" validate:"-" sql:"funding"`
	CompanyDetails Details `json:"company_details" validate:"-" sql:"companydetails"`
}

type Investor struct{
	Team
	InvestorType string `json:"investor_type" validate:"-" sql:"investortype"`
	InvestmentStage string `json:"investment_stage" validate:"-" sql:"investmentstage"`
	NumberOfExits int `json:"number_of_exits" validate:"-" sql:"numberofexits"`
	NumberOfinvestments int `json:"number_of_investments" validate:"-" sql:"numberofinvestments"`
	NumberOfFunds int `json:"number_of_funds" validate:"-" sql:"numberoffunds"`
}

type Funding struct {
	JsonEmbeddable
	TotalFunding int `json:"total_funding" validate:"required" sql:"totalfunding"`
	LatestRound string `json:"latest_round" validate:"required" sql:"latestround"`
	LatestRoundDate string `json:"latest_round_date" validate:"required" sql:"latestrounddate"`
	LatestRoundFunding int `json:"latest_round_funding" validate:"required" sql:"latestroundfunding"`
	Investors []string `json:"investors" validate:"-" sql:"investors"`
	InvestorType string `json:"investor_type" validate:"-" sql:"investortype"` //Accelerator
	InvestmentStage string `json:"investment_stage" validate:"-" sql:"investmentstage"` // Debt, Early Stage Venture, Seed
}

type Details struct {
	JsonEmbeddable
	IPOStatus string `json:"ipo_status" validate:"-" sql:"ipostatus"`
	CompanyType string `json:"company_type" validate:"-" sql:"companytype"` // for-profit
}

type TeamMember struct {
	JsonEmbeddable
	ID string `json:"ID" validate:"required" sql:"id"`
	Name string `json:"name" validate:"required" sql:"name"`
	Status string `json:"status" validate:"required" sql:"status"` // co-founder
}

type Contact struct {
	JsonEmbeddable
	Email string `json:"email" validate:"required" sql:"email"`
	PhoneNumber string `json:"phonenumber" validated:"required" sql:"phonenumber"`
}

type SocialMedia struct{
	JsonEmbeddable
	Website string `json:"website" validate:"-" sql:"website"`
	Facebook string `json:"facebook" validate:"-" sql:"facebook"`
	Twitter string `json:"twitter" validate:"-" sql:"twitter"`
	LinkedIn string `json:"linkedIn" validate:"-" sql:"linkedin"`
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