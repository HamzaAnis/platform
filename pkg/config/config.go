package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"SERVICE_PORT"`
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBUsername string `mapstructure:"POSTGRES_USERNAME"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DBNAME"`
}

var Cfg Config

func LoadConfig(path string) error {
	viper.AddConfigPath("./")

	viper.SetConfigName(path)
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	viper.AutomaticEnv()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Cfg)
	return err
}
