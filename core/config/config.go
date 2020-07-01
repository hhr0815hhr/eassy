package config

import (
	"github.com/spf13/viper"
	"os"
)

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
