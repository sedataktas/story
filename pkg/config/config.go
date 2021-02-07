package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/*Config is configuration object*/
var Config *Configuration

/*Configuration stores configuration fields. This struct sets from .yaml file*/
type Configuration struct {
	Database DatabaseConfiguration
}

// DatabaseConfiguration stores database credentials
type DatabaseConfiguration struct {
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
}

/*Setup sets configuration*/
func Setup() {
	viper.SetConfigName("story-conf-sample")
	viper.AddConfigPath("/Users/sedat/go/src/story")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var conf *Configuration
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	Config = conf
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
