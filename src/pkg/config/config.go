/*
 * @File: config.config.go
 * @Description: Defines common service configuration
 * @Author: Yoan Yomba (yoanyombapro@gmail.com)
 */
package config

import (
	"encoding/json"
	"os"
)

// Configuration stores setting values
type Configuration struct {
	Debug      string `json:"debug.addr"`
	Http       string `json:"http.addr"`
	Appdash       string `json:"appdash.addr"`
	DbType      string `json:"dbType"`
	DbAddress   string `json:"dbAddress"`
	DbName      string `json:"dbName"`
	DbSettings  string `json:"dbSettings"`
	Development bool   `json:"development"`
	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`
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

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}
	return nil
}

func (Config *Configuration) GetDatabaseConnectionString() string{
	return Config.DbType + Config.DbAddress + Config.DbName + Config.DbSettings
}