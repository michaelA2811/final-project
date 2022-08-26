package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(configFile string) {

	viper.SetConfigName(configFile)
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

}
