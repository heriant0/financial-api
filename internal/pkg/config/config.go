package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort              string        `mapstructure:"APP_PORT"`
	DatabaseURL          string        `mapstructure:"DATABASE_URL"`
	DatabaseDriver       string        `mapstructure:"DATABASE_DRIVER"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	HashCost             int           `mapstructure:"HASH_COST"`
	SecretKey            string        `mapstructure:"SECRET_KEY"`
	AccessTokenKey       string        `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey      string        `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
