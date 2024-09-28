package test

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitViper(t *testing.T) {
	viperNew := viper.New()
	viperNew.SetConfigName("config")
	viperNew.SetConfigType("json")
	viperNew.AddConfigPath("..")
	err := viperNew.ReadInConfig()

	assert.Nil(t, err)
}

func initViper() *viper.Viper {
	viperNew := viper.New()
	viperNew.SetConfigName("config")
	viperNew.SetConfigType("json")
	viperNew.AddConfigPath("..")
	err := viperNew.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return viperNew
}
