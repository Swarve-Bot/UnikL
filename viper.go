package unikl

import "github.com/spf13/viper"

func LoadConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	return viper.ReadInConfig()
}
