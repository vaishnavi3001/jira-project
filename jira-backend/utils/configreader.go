package utils

import (
	"fmt"
	ct "jira-backend/constants"
	"os"

	"github.com/spf13/viper"
)

var Vconfig *viper.Viper

func ConfigReader(env string) error {
	vInst := viper.New()
	vInst.SetConfigType("json")
	vInst.SetConfigName("project.config")
	vInst.AddConfigPath("$HOME/config")
	vInst.AddConfigPath("./config")
	if env == ct.TEST {
		vInst.AddConfigPath("../config")
	}
	vInst.ReadInConfig()
	Vconfig = vInst.Sub(ct.DEV)

	if env == ct.PROD {
		Vconfig = vInst.Sub(ct.PROD)
	}

	if Vconfig == nil {
		fmt.Println("Couldn't read config. Exiting")
		os.Exit(1)
	}

	return nil
}
