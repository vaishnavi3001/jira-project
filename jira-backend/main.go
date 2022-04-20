package main

import (
	"fmt"
	ct "jira-backend/constants"
	dt "jira-backend/dbutils"
	rt "jira-backend/routes"
	ut "jira-backend/utils"
	"os"
	"strings"
)

func main() {
	//Initialize The Configuration
	cmdArgs := os.Args
	env := ct.DEV
	if len(cmdArgs) > 1 {
		fmt.Println(cmdArgs[1])
		switch cmdArgs[1] {
		case ct.TEST:
			env = ct.TEST
		case ct.PROD:
			env = ct.PROD
		}
	}

	fmt.Printf("Environment Set to: %s\n", strings.ToUpper(env))

	err := ut.ConfigReader(env)
	if err != nil {
		fmt.Println("Cannot Read from the Config", err)
		os.Exit(1)
	}

	// Migrate the schema
	dberr := dt.InitializeConn(env)
	if dberr != nil {
		fmt.Println("Cannot Connect to the Database")
		os.Exit(1)
	}

	mailerr := ut.InitializeEmailSession()
	if mailerr != nil {
		fmt.Println("Cannot Connect to the Email Service")
		os.Exit(1)
	}

	r := rt.SetupRouter(true)
	ip_address := fmt.Sprintf("%s:%d", ut.Vconfig.GetString("server.ip_address"), ut.Vconfig.GetInt("server.port"))
	r.Run(ip_address)
}
