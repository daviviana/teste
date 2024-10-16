package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	JWTSecret         string `mapstructure:"JWT_SECRET"`
	JWTExpirationTime string `mapstructure:"JWT_EXPIRATION_TIME"`
	JWTIssuer         string `mapstructure:"DB_HOST"`
}

var AppConfig *Config

func LoadConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	AppConfig = &config
	return nil
}
