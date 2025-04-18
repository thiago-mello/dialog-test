package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	initializeViperConfig()
	initializeDatabase()
	initializeCache()
}

func initializeViperConfig() {
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configureDefaultEnvironmentVariables()

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

// configureDefaultEnvironmentVariables binds environment variables to configuration keys
func configureDefaultEnvironmentVariables() {
	// general envs
	viper.BindEnv("server.port", "PORT")
	viper.BindEnv("sql.templates.path", "SQL_TEMPLATES_PATH")
	viper.BindEnv("api.jwt.secret", "JWT_SECRET")
	viper.BindEnv("api.jwt.expires-in", "JWT_EXPIRES_IN")

	// database envs
	viper.BindEnv("database.relational.host", "DB_HOST")
	viper.BindEnv("database.relational.port", "DB_PORT")
	viper.BindEnv("database.relational.database-name", "DB_NAME")
	viper.BindEnv("database.relational.auth.user", "DB_USER")
	viper.BindEnv("database.relational.auth.password", "DB_PASSWORD")

	// redis envs
	viper.BindEnv("database.redis.address", "REDIS_ADDR")
	viper.BindEnv("database.redis.password", "REDIS_PASSWORD")
	viper.BindEnv("database.redis.db", "REDIS_DB")

	// instrumentation
	viper.BindEnv("otel.traces.otlp.endpoint", "TRACES_OTLP_ENDPOINT")
}
