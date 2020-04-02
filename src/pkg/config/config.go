/*
 * @File: config.config.go
 * @Description: Defines common service configuration
 * @Author: Yoan Yomba (yoanyombapro@gmail.com)
 */
package config

// Configuration stores setting values
type Configuration struct {
	ServerConfiguration
	AmqpConfiguration
}

// ServerConfiguration witholds important parameters such as ports and service
// names
type ServerConfiguration struct {
	Port           string `arg:"env:SERVER_PORT"`
	Name           string `arg:"env:SERVICE_NAME`
	Debug          string `arg:"env:DEBUG_ADDR"`
	Http           string `arg:"env:HTTP_ADDR"`
	Appdash        string `arg:"env:APPDASH_ADDR"`
	ZipkinUrl      string `arg:"env:ZIPKIN_URL"`
	UseZipkin      bool   `arg:"env:ZIPKIN_USE"`
	Zipkin         string `arg:"env:ZIPKIN_ADDR"`
	DbType         string `arg:"env:DB_TYPE"`
	DbAddress      string `arg:"env:DB_ADDRESS"`
	DbName         string `arg:"env:DB_NAME"`
	DbSettings     string `arg:"env:DB_SETTINGS"`
	Development    bool   `arg:"env:DEVELOPMENT"`
	Jwt            string `arg:"env:JWTSECRETPASSWORD"`
	Issuer         string `arg:"env:ISSUER"`
	ZipkinBridge   bool   `arg:"env:ZIPKINBRIDGE"`
	LightstepToken string `arg:"env:LIGHTSTEP"`
}

// AmqpConfiguration witholds connections parameters for a
// rabbitMQ instance
type AmqpConfiguration struct {
	ServerUrl string `arg:"env:AMQP_SERVER_URL"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

// Status Text
const (
	ErrNameEmpty      = "Name is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

// Status Code
const (
	StatusCodeUnknown = -1
	StatusCodeOK      = 1000
)

// DefaultConfiguration provides a default configuration object
func DefaultConfiguration() {
	Config = &Configuration{
		ServerConfiguration: ServerConfiguration{
			Name:           "users_microservice",
			Port:           "6868",
			Issuer:         "cubeplatform",
			Jwt:            "cubeplatformjwtpassword",
			Development:    true,
			DbSettings:     "?sslmode=require",
			DbName:         "defaultdb",
			DbAddress:      "doadmin:oqshd3sto72yyhgq@test-do-user-6612421-0.a.db.ondigitalocean.com:25060/",
			DbType:         "postgresql://",
			Zipkin:         ":8080",
			UseZipkin:      true,
			ZipkinUrl:      "http://localhost:9411/api/v2/spans",
			Appdash:        ":8086",
			Http:           ":8085",
			Debug:          ":8084",
			ZipkinBridge:   true,
			LightstepToken: "",
		},
		AmqpConfiguration: AmqpConfiguration{
			ServerUrl: "amqp://guest:guest@rabbitmq:5672/",
		},
	}
}

// GetDatabaseConnectionString Creates a database connection string from the service configuration settings
func (Config *Configuration) GetDatabaseConnectionString() string {
	if Config == nil {
		DefaultConfiguration()
	}
	return Config.ServerConfiguration.DbType + Config.ServerConfiguration.DbAddress +
		Config.ServerConfiguration.DbName + Config.ServerConfiguration.DbSettings
}
