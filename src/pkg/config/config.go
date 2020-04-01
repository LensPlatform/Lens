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
	Environment    string `arg:"Environment"`
	Port           string `arg:"env:SERVER_PORT"`
	Name           string `arg:"env:SERVICE_NAME`
	Debug          string `arg:"debug_addr"`
	Http           string `arg:"env:http_addr"`
	Appdash        string `arg:"env:appdash_addr"`
	ZipkinUrl      string `arg:"env:zipkin_url"`
	UseZipkin      bool   `arg:"env:zipkin_use"`
	Zipkin         string `arg:"env:zipkin_addr"`
	DbType         string `arg:"env:dbType"`
	DbAddress      string `arg:"env:dbAddress"`
	DbName         string `arg:"env:dbName"`
	DbSettings     string `arg:"env:dbSettings"`
	Development    bool   `arg:"env:development"`
	Jwt            string `arg:"env:jwtSecretPassword"`
	Issuer         string `arg:"env:issuer"`
	ServiceName    string `arg:"env:serviceName"`
	ZipkinBridge   bool   `arg:"env:zipkin-ot-bridge"`
	LightstepToken string `arg:"env:lightstep-token"`
}

// AmqpConfiguration witholds connections parameters for a
// rabbitMQ instance
type AmqpConfiguration struct {
	ServerUrl string `arg:"env.AMQP_SERVER_URL"`
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
			Environment:    "development",
			Name:           "users_microservice",
			Port:           "6868",
			ServiceName:    "users_microservice",
			Issuer:         "cubeplatform",
			Jwt:            "lensplatformjwtpassword",
			Development:    false,
			DbSettings:     "?sslmode=require",
			DbName:         "users-microservice-db",
			DbAddress:      "doadmin:x9nec6ffkm1i3187@backend-datastore-do-user-6612421-0.db.ondigitalocean.com:25060/",
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
