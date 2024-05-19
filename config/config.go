package config

import (
	"github.com/spf13/viper"
)

type CryptoConfig struct {
	HS256Password            string
	JWSPublicKey             string // ED-25519
	JWSPrivateKey            string // ED-25519
	JWSPrivateKeyPassword    string
	JWSExpirationTimeInHours int64
}

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
	CryptoConfig   *CryptoConfig
}

func Get() *Config {
	viper.AutomaticEnv()

	return &Config{
		CryptoConfig: &CryptoConfig{
			HS256Password:            viper.GetString("HS256_PASSWORD"),
			JWSPublicKey:             viper.GetString("JWS_PUBLIC_KEY"),
			JWSPrivateKey:            viper.GetString("JWS_PRIVATE_KEY"),
			JWSPrivateKeyPassword:    viper.GetString("JWS_PRIVATE_KEY_PASSWORD"),
			JWSExpirationTimeInHours: viper.GetInt64("JWS_EXPIRATION_TIME_IN_HOURS"),
		},
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
