package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mode string `mapstructure:"mode"`

	Server struct {
		Name string `mapstructure:"name"`
		Http struct {
			Address string `mapstructure:"address"`
			Prefix  string `mapstructure:"prefix"`
		} `mapstructure:"http"`
	} `mapstructure:"server"`

	Mongo struct {
		Uri string `mapstructure:"uri"`
		DB  string `mapstructure:"db"`
	}
}

var common *Config

func Get() *Config {
	return common
}

func LoadConfig(pathConfig string) error {
	viper.SetConfigFile(pathConfig)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&common)

	return nil
}
