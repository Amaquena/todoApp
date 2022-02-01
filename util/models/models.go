package models

type DBConfig struct {
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Hostname string `mapstructure:"DB_HOSTNAME"`
	Port     string `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_DBNAME"`
}