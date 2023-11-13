package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Load() (bool, error) {
	runPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	if _, err := os.Stat(runPath + "/" + "credentials.yml"); os.IsNotExist(err) {
		fmt.Println("credentials.yml doesn't exists")
		os.Exit(1)
		return false, err
	}

	viper.SetConfigFile(runPath + "/" + CONFIG_PATH)
	err = viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
