package database

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	model "github.com/LensPlatform/Lens/src/pkg/models"
	_ "github.com/LensPlatform/Lens/src/pkg/helper"
)

type DBHandler interface {
	// Create
	CreateUser(user model.User) error
	CreateTeam(founder model.User, team model.Team) error

	// GET
	GetUserById(id string) (model.Team, error)
	GetUserByEmail(email string) (model.Team, error)
	GetUserByUsername(username string) (model.Team, error)
	GetPassword(id string)( string,  error)
	GetAllUsers()([]model.User, error)
	GetAllUsersFromSearchQuery(search map[string]interface{}) ([]model.User, error)
	GetUserBasedOnParam(param string, query string) (model.User, error)

	GetTeamByID(id string) (model.Team, error)
	GetTeamByName(name string) (model.Team, error)
	GetTeamsByType(teamType string) ([]model.Team, error)
	GetTeamsByIndustry(industry string) ([]model.Team, error)
	GetAllTeams() ([]model.Team, error)
	GetAllTeamsFromSearchQuery(search map[string]interface{}) ([]model.Team, error)
	GetTeamBasedOnParam(param string, query string) (model.Team, error)

	// Update
	UpdateUser(param map[string]interface{}, id string) (model.Team, error)

	UpdateTeamName(name string, teamId string) (model.Team, error)
	UpdateTeamType(teamType string, teamId string) (model.Team, error)
	UpdateTeamOverview(overView string, teamId string) (model.Team, error)
	updateTeamIndustry(industry string, teamId string) (model.Team, error)
	AddTeamMemberToTeam(teamMember model.TeamMember, teamId string) (model.Team, error)
	RemoveTeamMemberFromTeam(teamMember model.TeamMember, teamId string) (model.Team, error)
	AddAdvisorToTeam(advisorMember model.TeamMember, teamId string) (model.Team, error)
	RemoveAdvisorFromTeam(advisorMember model.TeamMember, teamId string)(model.Team, error)
	AddFounderToTeam(advisorMember model.TeamMember, teamId string) (model.Team, error)
	RemoveFounderFromTeam(advisorMember model.TeamMember, teamId string) (model.Team, error)

	// Delete
	DeleteUserById(id string) (bool, error)
	DeleteUserByUsername(id string) (bool, error)
	DeleteUserByEmail(id string) (bool, error)
	DeleteUserBasedOnParam(param string, query string) (bool, error)

	DeleteTeamById(teamId string) (bool, error)
	DeleteTeamByName(teamName string) (bool, error)
	DeleteTeamTeamByEmail(teamEmail string) (bool, error)
	DeleteTeamBasedOnParam(param string, query string) (bool, error)

	// Existence
	DoesUserExist(searchParam string, query string) (bool, error)
	DoesTeamExist(searchParam string, query string) (bool, error)
}

type Database struct {
	connection *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{connection:db}
}

func (db Database) CreateUser(user model.User) error {
	if user.ID == ""{
		errMsg := fmt.Sprintf("Invalid Argument provided. " +
			"the following param is null User Id : %s", user.ID)
		return errors.New(errMsg)
	}

	e := db.connection.Create(&user).Error

	if e != nil {
		return e
	}
	return nil
}

func (db Database) CreateTeam(founder model.User, team model.Team) error {
	if founder.ID == ""  || team.ID == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of " +
			"the following params are null Team Id : %s, Founder ID : %s", founder.ID, team.ID)
		return errors.New(errMsg)
	}

	var teamMember model.TeamMember
	var founders []model.TeamMember

	founderName := fmt.Sprintf("%s %s", founder.FirstName, founder.LastName)
	teamMember = model.TeamMember{ID : founder.ID, Name : founderName, Title: "founder"}

	founders = append(team.Founders, teamMember)
	team.Founders = founders
	team.NumberOfEmployees = len(team.TeamMembers) + len(team.Advisors) + len(team.Founders)

	e := db.connection.Create(&team).Error

	if e != nil {
		return e
	}
	return nil
}

func (db Database) GetUserById(id string) (model.User, error) {
	query := "id = ?"
	return db.GetUserBasedOnParam(id, query)
}

func (db Database) GetUserByUsername(username string) (model.User, error)  {
	query := "username = ?"
	return db.GetUserBasedOnParam(username,query)
}

func (db Database) GetUserByEmail(email string) (model.User, error)  {
	query := "email = ?"
	return db.GetUserBasedOnParam(email,query)
}

func (db Database) GetAllUsers() ([]model.User, error) {
	var users []model.User
	e := db.connection.Find(&users).Error
	if e != nil {
		return nil, e
	}
	return users, nil
}
func (db Database) GetAllUsersFromSearchQuery(query map[string]interface{}) ([]model.User, error) {
	var users []model.User
	// ex. db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	e := db.connection.Where(query).Find(&users).Error
	if e != nil {
		return nil, e
	}
	return users, nil
}

func (db Database) GetUserBasedOnParam(param string, query string) (model.User, error) {
		if param == ""  || query == "" {
			errMsg := fmt.Sprintf("Invalid Argument provided. One of " +
				"the following params are null Search Param : %s, Query : %s", param, query)
			return model.User{}, errors.New(errMsg)
		}
		var user model.User
		e := db.connection.First(&user, query, param).Error

		if e != nil{
			return model.User{}, e
		}
		return user, e
}

func (db Database) GetTeamByID(id string) (model.Team, error) {
	if id == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. " +
			"The following argument is null Id : %s", id)
		return model.Team{}, errors.New(errMsg)
	}
	query := "id = ?"
	return db.GetTeamBasedOnParam(id, query)
}

func (db Database) GetTeamByName(name string) (model.Team, error) {
	if name == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. " +
			"The following argument is null name : %s", name)
		return model.Team{}, errors.New(errMsg)
	}
	query := "name = ?"
	return db.GetTeamBasedOnParam(name, query)
}

func (db Database) GetTeamsByType(teamType string) (model.Team, error) {
	if teamType == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. " +
			"The following argument is null Team Type : %s", teamType)
		return model.Team{}, errors.New(errMsg)
	}
	query := "type = ?"
	return db.GetTeamBasedOnParam(teamType, query)
}

func (db Database) GetTeamsByIndustry(industry string) ([]model.Team, error) {
	if industry == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. " +
			"The following argument is null Industry Type : %s", industry)
		return nil, errors.New(errMsg)
	}
	query := "industry = ?"
	var teams []model.Team
	e := db.connection.Where(query, industry).Find(&teams).Error

	if e != nil {
		return nil, e
	}
	return teams, nil
}

func (db Database) GetAllTeams()([]model.Team, error){
	var teams []model.Team
	e := db.connection.Find(&teams).Error

	if e != nil {
		return  nil, e
	}
	return teams, nil
}

func (db Database) GetAllTeamsFromSearchQuery(search map[string]interface{}) ([]model.Team , error){
	var teams []model.Team
	e := db.connection.Where(search).Find(&teams).Error

	if e != nil {
		return  nil, e
	}
	return teams, nil
}

func (db Database) GetTeamBasedOnParam(param string, query string)(model.Team, error){
	if param == ""  || query == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of " +
			"the following params are null Search Param : %s, Query : %s", param, query)
		return model.Team{}, errors.New(errMsg)
	}

	var team model.Team
	e := db.connection.First(&team, query, param).Error

	if e != nil{
		return model.Team{}, e
	}
	return team, nil
}

func (db Database) DoesUserExist(searchParam string, query string) (bool, error) {
	// check if user exists
	var user model.User
	user, err := db.GetUserBasedOnParam(searchParam,query)

	if err != nil {
		return false, err
	}

	if user.ID != ""{
		return true, nil
	}

	return false, nil
}

func (db Database) DoesTeamExist(searchParam string, query string) (bool, error) {
	var team model.Team
	team, err := db.GetTeamBasedOnParam(searchParam,query)

	if err != nil {
		return false, err
	}

	if team.ID != "" {
		return true, nil
	}

	return false, nil
}

func (db Database) UpdateUser(param map[string]interface{}, id string) (model.User, error) {
	var user model.User
	user, err := db.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	// db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	err = db.connection.Model(&user).Updates(param).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (db Database) UpdateTeamName(name string, teamId string) (model.Team, error) {
	var team model.Team
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	err = db.connection.Model(&team).UpdateColumn("name", name).Error
	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) UpdateTeamType(teamType string, teamId string) (model.Team, error) {
	var team model.Team
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	err = db.connection.Model(&team).UpdateColumn("type", teamType).Error

	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) UpdateTeamOverview(overView string, teamId string) (model.Team, error) {
	var team model.Team
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	err = db.connection.Model(&team).UpdateColumn("overview", overView).Error
	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) UpdateTeamIndustry(industry string, teamId string)  (model.Team, error) {
	var team model.Team
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	err = db.connection.Model(&team).UpdateColumn("industry", industry).Error
	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) AddTeamMemberToTeam(teamMember model.TeamMember, teamId string) (model.Team, error) {
	if teamMember.ID == ""  || teamId == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of the following params are null Team Id : %s, Team Member" +
			" ID : %s", teamMember.ID, teamId)
		return model.Team{}, errors.New(errMsg)
	}

	var team model.Team
	var teamMembers []model.TeamMember
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	teamMembers = append(team.TeamMembers, teamMember)
	team.TeamMembers = teamMembers
	team.NumberOfEmployees = len(team.TeamMembers) + len(team.Advisors) + len(team.Founders)

	err = db.connection.Save(&team).Error

	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) RemoveTeamMemberFromTeam(teamMember model.TeamMember, teamId string) (model.Team, error) {
	if teamMember.ID == ""  || teamId == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of the following params are null Team Id : %s, Team Member" +
			" ID : %s", teamMember.ID, teamId)
		return model.Team{}, errors.New(errMsg)
	}

	var team model.Team
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	var index int

	for i, item := range team.TeamMembers {
		if item.ID == teamMember.ID {
			index = i
			break
		}
	}

	team.TeamMembers[index] = team.TeamMembers[len(team.TeamMembers)-1]
	team.TeamMembers = team.TeamMembers[:len(team.TeamMembers)-1]

	team.NumberOfEmployees = len(team.TeamMembers) + len(team.Advisors) + len(team.Founders)
	err = db.connection.Save(&team).Error

	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) AddAdvisorToTeam(advisorMember model.TeamMember, teamId string) (model.Team, error) {
	if advisorMember.ID == ""  || teamId == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of the following params are null Team Id : %s, Team Member" +
			" ID : %s", advisorMember.ID, teamId)
		return model.Team{}, errors.New(errMsg)
	}

	var team model.Team
	var advisorMembers []model.TeamMember
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	advisorMembers = append(team.Advisors, advisorMember)
	team.Advisors = advisorMembers
	team.NumberOfEmployees = len(team.TeamMembers) + len(team.Advisors) + len(team.Founders)

	err = db.connection.Save(&team).Error

	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) RemoveAdvisorFromTeam(advisorMember model.TeamMember, teamId string) (model.Team, error) {
	if advisorMember.ID == ""  || teamId == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of the following params are null Team Id : %s, Team Member" +
			" ID : %s", advisorMember.ID, teamId)
		return model.Team{}, errors.New(errMsg)
	}

	var team model.Team
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	var index int

	for i, item := range team.Advisors {
		if item.ID == advisorMember.ID {
			index = i
			break
		}
	}

	team.Advisors[index] = team.Advisors[len(team.Advisors)-1]
	team.Advisors = team.Advisors[:len(team.Advisors)-1]

	team.NumberOfEmployees = len(team.TeamMembers) + len(team.Advisors) + len(team.Founders)
	err = db.connection.Save(&team).Error

	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) AddFounderToTeam(founderMember model.TeamMember, teamId string)  (model.Team, error) {
	if founderMember.ID == ""  || teamId == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of the following params are null Team Id : %s, Team Member" +
			" ID : %s", founderMember.ID, teamId)
		return model.Team{}, errors.New(errMsg)
	}

	var team model.Team
	var founderMembers []model.TeamMember
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	founderMembers = append(team.Founders, founderMember)
	team.Advisors = founderMembers
	team.NumberOfEmployees = len(team.TeamMembers) + len(team.Advisors) + len(team.Founders)

	err = db.connection.Save(&team).Error

	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) RemoveFounderFromTeam(founderMember model.TeamMember, teamId string) (model.Team, error) {
	if founderMember.ID == ""  || teamId == "" {
		errMsg := fmt.Sprintf("Invalid Argument provided. One of the following params are null Team Id : %s, Team Member" +
			" ID : %s", founderMember.ID, teamId)
		return model.Team{}, errors.New(errMsg)
	}

	var team model.Team
	team, err := db.GetTeamByID(teamId)

	if err != nil {
		return model.Team{}, err
	}

	var index int

	for i, item := range team.Advisors {
		if item.ID == founderMember.ID {
			index = i
			break
		}
	}

	team.Founders[index] = team.Founders[len(team.Founders)-1]
	team.Founders = team.Founders[:len(team.Founders)-1]

	team.NumberOfEmployees = len(team.TeamMembers) + len(team.Advisors) + len(team.Founders)

	err = db.connection.Save(&team).Error

	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (db Database) DeleteUserById(id string) (bool, error){
	query := "id = ?"
	return db.DeleteUserBasedOnParam(id,query)
}

func (db Database) DeleteUserByUsername(id string) (bool, error){
	query := "username = ?"
	return db.DeleteUserBasedOnParam(id,query)
}

func (db Database) DeleteUserByEmail(id string) (bool, error){
	query := "email = ?"
	return db.DeleteUserBasedOnParam(id,query)
}

func (db Database) 	DeleteUserBasedOnParam(param string, query string) (bool, error) {
	var user model.User
	user, err := db.GetUserBasedOnParam(param,query)

	if err != nil {
		return false, err
	}

	if user.ID == ""{
		return false, nil
	}

	err = db.connection.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&user).Error

	if err != nil {
		return false, nil
	}
	return true, nil
}

func (db Database) DeleteTeamById(teamId string) (bool, error) {
	query := "id = ?"
	return db.DeleteTeamBasedOnParam(teamId,query)
}

func (db Database) DeleteTeamByName(teamName string) (bool, error) {
	query := "name = ?"
	return db.DeleteTeamBasedOnParam(teamName,query)
}

func (db Database) DeleteTeamTeamByEmail(teamEmail string) (bool, error){
	query := "email = ?"
	return db.DeleteTeamBasedOnParam(teamEmail,query)
}

func (db Database) 	DeleteTeamBasedOnParam(param string, query string) (bool, error) {
	var team model.Team
	team, err := db.GetTeamBasedOnParam(param,query)

	if err != nil {
		return false, err
	}

	if team.ID == ""{
		return false, nil
	}

	err = db.connection.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&team).Error
	if err != nil {
		return false, err
	}
	return true, nil
}