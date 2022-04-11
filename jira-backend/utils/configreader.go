package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

var Vconfig *viper.Viper

func ConfigReader(isTest bool) error {
	Vconfig = viper.New()
	Vconfig.SetConfigType("json")
	Vconfig.SetConfigName("project.config")
	Vconfig.AddConfigPath("$HOME/jira-clone/")
	Vconfig.AddConfigPath("./config")

	if isTest {
		Vconfig.AddConfigPath("../config")
	}

	fmt.Println(Vconfig.GetInt("server.port"))
	err := Vconfig.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
