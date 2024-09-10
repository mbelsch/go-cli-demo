package config

import (
	"fmt"
	"mbelsch/helper/pkg/utils"
	"os"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

// GetConfig gets the configuration for the given environment.
// It returns the configuration and an error if one occured.
func GetConfig(env string) (*Config, error) {
	err := CheckConfigExists(env)
	if err != nil {
		return nil, err
	}

	configPath, err := GetConfigPath(env)
	if err != nil {
		return nil, err
	}

	return GetConfigFromFile(configPath)
}

// CheckConfigExists tests if the environment configuration file exists and returns an error if not.
func CheckConfigExists(env string) error {
	path, err := GetConfigPath(env)
	if err != nil {
		return err
	}

	exists, err := utils.DoesFileExist(path)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the configuration file for env %s does not exist on path: %s", env, path)
	}

	return nil
}

// GetConfigPath returns the path of the config file for the given environment.
func GetConfigPath(env string) (path string, err error) {
	path, err = GetConfigDirectory()
	if err != nil {
		return "", err
	}

	path += "/" + env + ".yaml"

	return path, nil
}

// GetConfigDirectory returns the path to the .helper folder containing the configuration.
// By default it is the folder ~/.helper.
// However when running on the CI server it is the /root/.helper as build servers tend to overwrite
// home folder location and our docker container puts all the files in /root/.helper.
func GetConfigDirectory() (path string, err error) {
	// on github the CI variable is always true
	if os.Getenv("CI") != "" {
		path := "/root/.helper"
		fmt.Println("CI variable is set, using " + path + " as config dir")
		return path, nil
	}

	path, err = homedir.Dir()
	if err != nil {
		return "", err
	}

	path += "/.helper"

	return path, nil
}

// GetConfigFromFile gets the configuration for the given file path.
// It returns the configuration and an error if one occured.
func GetConfigFromFile(configFilePath string) (*Config, error) {
	//reads config and logs error in case
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf(err.Error() + " Error while reading config from project root")
		return nil, err
	}

	var config Config

	//parses read config into struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf(err.Error() + " Error while reading config")
		return nil, err
	}

	return &config, nil
}
