package main

import (
	"fmt"
	dt "jira-backend/dbutils"
	"jira-backend/routes"
	ut "jira-backend/utils"
	"os"
)

func main() {
	//Initialize The Configuration
	err := ut.ConfigReader(false)
	if err != nil {
		fmt.Println("Cannot Read from the Config", err)
		os.Exit(1)
	}

	// Migrate the schema
	dberr := dt.InitializeConn(false)
	if dberr != nil {
		fmt.Println("Cannot Connect to the Database")
		os.Exit(1)
	}

	mailerr := ut.InitializeEmailSession()
	if mailerr != nil {
		fmt.Println("Cannot Connect to the Email Service")
		os.Exit(1)
	}
	r := routes.SetupRouter(true)
	ip_address := fmt.Sprintf("%s:%d", ut.Vconfig.GetString("server.ip_address"), ut.Vconfig.GetInt("server.port"))
	r.Run(ip_address)
}
