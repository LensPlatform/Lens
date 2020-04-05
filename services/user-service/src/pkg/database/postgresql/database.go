package postgresql

import (
	"os"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"github.com/LensPlatform/Lens/services/user-service/src/pkg/config"
	table "github.com/LensPlatform/Lens/services/user-service/src/pkg/models/proto"
)

type IDatabase interface {
	CreateUser(User table.UserORM) error
	UpdateUser(User table.UserORM) error
	DeleteUser(User table.UserORM) error
	GetUserById(id int32) (error, *table.UserORM)
	GetUserByUsername(username string) (error, *table.UserORM)
	GetUserByEmail(email string) (error, *table.UserORM)
	GetAllUsers(limit int) (error, []*table.UserORM)

	CreateGroup(group table.GroupORM) error
	UpdateGroup(group table.GroupORM) error
	DeleteGroup(group table.GroupORM) error
	GetGroupById(id int32) (error, *table.GroupORM)
	GetGroupByName(name string) (error, *table.GroupORM)
	GetAllGroups(limit int) (error, []*table.GroupORM)

	CreateTeam(group table.TeamORM) error
	UpdateTeam(group table.TeamORM) error
	DeleteTeam(group table.TeamORM) error
	GetTeamById(id int32) (error, *table.TeamORM)
	GetTeamByName(name string) (error, *table.TeamORM)
	GetAllTeams(limit int) (error, []*table.TeamORM)
}

type Database struct {
	Engine *gorm.DB
	Logger *zap.Logger
}

var (
	postgres = "postgres"
)

func New(conn string, logger *zap.Logger) (error, *Database) {
	db, err := gorm.Open(postgres, conn)
	if err != nil {
		logger.Error(err.Error())
		return err, nil
	}

	logger.Info("Successfully connected to the database")

	db.SingularTable(true)
	db.LogMode(false)

	logger.Info("Auto Migrating database tables")

	db.AutoMigrate(table.TeamORM{})
	db.AutoMigrate(table.UserORM{})
	db.AutoMigrate(table.GroupORM{})
	db.AutoMigrate(table.AddressORM{})

	logger.Info("Auto Migration of database tables complete")

	return nil, &Database{
		Engine: db,
		Logger: logger,
	}
}

// InitDbConnection initializes a database connection and creates associated tables/migrates schemas
func Init(zapLogger *zap.Logger) (*gorm.DB, error) {
	connString := config.Config.GetDatabaseConnectionString()
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		zapLogger.Error(err.Error())
		os.Exit(1)
	}
	// db.Set("gorm:table_options", "ENGINE=InnoDB")
	zapLogger.Info("successfully connected to database")
	db.SingularTable(true)
	db.LogMode(false)
	CreateTablesOrMigrateSchemas(db, zapLogger)

	return db, err
}


// CreateTablesOrMigrateSchemas creates a given set of tables based on a schema
// if it does not exist or migrates the table schemas to the latest version
func CreateTablesOrMigrateSchemas(db *gorm.DB, zapLogger *zap.Logger) {
	db.AutoMigrate(table.AddressORM{}, table.EducationORM{},table.MediaORM{}, table.SubscriptionsORM{}, table.SocialMediaORM{},
	table.DetailsORM{}, table.ExperienceORM{}, table.InvestmentORM{}, table.UserORM{}, table.ProfileORM{}, table.GroupORM{},
	table.TeamORM{},table.TeamProfileORM{}, table.InvestorDetailORM{}, table.StartupDetailORM{}, table.SettingsORM{}, table.LoginActivityORM{},
	table.PaymentsORM{}, table.CardORM{}, table.PinORM{}, table.Privacy{},table.NotificationORM{}, table.PostAndCommentsPushNotificationORM{},
	table.FollowingAndFollowersPushNotificationORM{}, table.DirectMessagesPushNotificationORM{}, table.EmailAndSmsPushNotificationORM{})
}
