package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mode string `mapstructure:"mode"`

	Cors struct {
		EnableLocal bool   `mapstructure:"enable_local"`
		Localhost   string `mapstructure:"localhost"`
	} `mapstructure:"cors"`

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

	Redis struct {
		Hosts    []string `mapstructure:"hosts"`
		Username string   `mapstructure:"username"`
		Password string   `mapstructure:"password"`
	} `mapstructure:"redis"`

	MQTT struct {
		Name     string `mapstructure:"name"`
		ClientId string `mapstructure:"client_id"`
		Host     string `mapstructure:"host"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mqtt"`
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
