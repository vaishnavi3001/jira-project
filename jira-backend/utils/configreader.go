package utils

import (
	"github.com/spf13/viper"
)

func ConfigReader() (*viper.Viper, error) {

	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigName("project.config")

	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	//fmt.Println(v.ConfigFileUsed())
	//fmt.Println(v.GetString("server.ip_address"))

	return v, nil
}
