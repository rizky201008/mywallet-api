package app

import "github.com/spf13/viper"

var Vipers *viper.Viper

func InitViper() {
	viperNew := viper.New()
	viperNew.SetConfigName("config")
	viperNew.SetConfigType("json")
	viperNew.AddConfigPath(".")
	err := viperNew.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Vipers = viperNew
}
