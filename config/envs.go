package config

import (
	"fmt"
	"os"
)

var (
	env string
)

// InitEnvs loads the environment variables we use to configure the rest of the app
func InitEnvs() {
	env := os.Getenv("ENV")
	if env != "dev" && env != "prod" {
		fmt.Printf("$DEV not set to either 'dev' or 'prod'")
		os.Exit(1)
	}
}

// IsDev returns whether or not the environment we are currently running in is a dev environment
func IsDev() bool {
	return env == "dev"
}

// IsProd returns whether or not the environment we are currently running in is a prod environment
func IsProd() bool {
	return env == "prod"
}
