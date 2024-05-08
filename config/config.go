package config

import (
	"github.com/spf13/viper"
)

type InternalConfig struct {
	RunningLocal bool
	ServerPort   int
	ServiceName  string
}

type MySQLConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Database int
}

type Config struct {
	InternalConfig *InternalConfig
	MySQLConfig    *MySQLConfig
	RedisConfig    *RedisConfig
}

func Get() *Config {
	viper.AutomaticEnv()

	return &Config{
		InternalConfig: &InternalConfig{
			RunningLocal: viper.GetBool("RUNNING_LOCAL"),
			ServerPort:   viper.GetInt("SERVER_PORT"),
			ServiceName:  viper.GetString("SERVICE_NAME"),
		},
		MySQLConfig: &MySQLConfig{
			Host:     viper.GetString("MYSQL_HOST"),
			Port:     viper.GetString("MYSQL_PORT"),
			Username: viper.GetString("MYSQL_USERNAME"),
			Password: viper.GetString("MYSQL_PASSWORD"),
			Database: viper.GetString("MYSQL_DATABASE"),
		},
		RedisConfig: &RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetString("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			Database: viper.GetInt("REDIS_DATABASE"),
		},
	}
}
