package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
			return
		}

		log.Fatal("error reading config file")
	}
}

// Merges new config file to the existing configuration. Usefull for tests
func MergeNewConfig(configPaths []string) {
	for _, path := range configPaths {
		viper.AddConfigPath(path)
		err := viper.MergeInConfig()
		if err != nil {
			fmt.Println("config file not found")
		}
	}
}

// Get config variable
func Get(key string) any {
	return viper.Get(key)
}

// Get string config variable
func GetString(key string) string {
	return viper.GetString(key)
}

// Get integer config variable
func GetInt(key string) int {
	return viper.GetInt(key)
}

// Get a String slice config variable
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// Get a boolean config variable
func GetBoolean(key string) bool {
	return viper.GetBool(key)
}
