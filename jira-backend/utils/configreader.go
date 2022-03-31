package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var Vconfig *viper.Viper

func ConfigReader() error {
	Vconfig = viper.New()
	Vconfig.SetConfigType("json")
	Vconfig.SetConfigName("project.config")

	Vconfig.AddConfigPath("./config")
	fmt.Println(Vconfig.GetInt("server.port"))
	err := Vconfig.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
