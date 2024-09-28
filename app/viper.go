package app

import "github.com/spf13/viper"

func InitViper() *viper.Viper {
	viperNew := viper.New()
	viperNew.SetConfigName("config")
	viperNew.SetConfigType("json")
	viperNew.AddConfigPath(".")
	err := viperNew.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viperNew
}
