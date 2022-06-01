package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	HostName string `mapstructure:"hostname"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"dbName"`
}

type Application struct {
	LogLevel string `mapstructure:"logLevel"`
	// if MockStorage is true we would use SQLite access temporary system memory, else we'll use MySQL
	MockStorage    bool  `mapstructure:"mockStorage"`
	RequestTimeOut int64 `mapstructure:"requestTimeout"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Config struct {
	Database    Database    `mapstructure:"database"`
	Server      Server      `mapstructure:"server"`
	Application Application `mapstructure:"application"`
}

func NewConfig() (*Config, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file %s", err)
	}

	var conf Config

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("error decoding config values %s", err)
	}

	return &conf, nil
}
