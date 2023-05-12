package server

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Url      string
	Password string
	Username string
}

func getConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s", err)
		return nil, err
	}

	return &Config{Url: viper.GetString("REDIS_URL"), Password: viper.GetString("REDIS_PASSWORD"), Username: viper.GetString("REDIS_USERNAME")}, nil
}
