package cmd

import (
	"fmt"
	"mbelsch/helper/pkg/config"
	"mbelsch/helper/pkg/exec"
	"mbelsch/helper/pkg/utils"
	"os"

	"github.com/spf13/cobra"
)

type Login struct {
	env string
}

var loginParams = Login{}

func initLogin(parent *cobra.Command) {
	parent.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&loginParams.env, "env", "e", "", "Region to login to")
	loginCmd.MarkFlagRequired("env")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login all services to specified environment",
	Long:  `login all services to specified environment`,
	Run: func(cmd *cobra.Command, args []string) {
		env := loginParams.env
		config, err := config.GetConfig(env)
		checkErrorAndExit(err, "get config for env %s", env)

		err = loginAz(config)
		checkErrorAndExit(err, "login az")
	},
}

//helper functions

func checkErrorAndExit(err error, action string, params ...string) {
	if err != nil {
		fmt.Println(err, "Error on: "+action, params)
		os.Exit(1)
	}
}

func loginAz(config *config.Config) error {

	//Setting Cloud for specific region
	err := exec.ExecCmd("az", "cloud", "set", "--name", config.Azure.CloudName)
	if err != nil {
		return err
	}

	var auth = ""
	if config.Azure.CertPath == "" && config.Azure.Password == "" {
		fmt.Printf("Neither certificate path nor password for service principle was provided. Login failed")
		os.Exit(1)
	} else if config.Azure.CertPath != "" {
		auth = config.Azure.CertPath
		existCerts, err := utils.DoesFileExist(auth)
		if err != nil {
			return err
		}
		if !existCerts {
			err = fmt.Errorf("client certificates not found under path %s", auth)
		}
		if err != nil {
			return err
		}
	} else {
		auth = config.Azure.Password
	}

	//login
	err = exec.ExecCmd("az",
		"login", "--service-principal", "-u", config.Azure.ServicePrincipal.AppID, "-p", auth, "-t", config.Azure.TenantID)
	if err != nil {
		return err
	}

	//Set SubscriptionID
	_, _, err = exec.ExecCmdOut("az", "account", "set", "--subscription", config.Azure.SubscriptionID)
	if err != nil {
		return err
	}
	return nil
}
