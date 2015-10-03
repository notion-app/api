package config

import (
	"fmt"
	"os"
)

var (
	env  string
	port string
)

// InitEnvs loads the environment variables we use to configure the rest of the app
func InitEnvs() {
	env = os.Getenv("ENV")
	if env != "dev" && env != "prod" {
		fmt.Printf("$DEV not set to either 'dev' or 'prod'")
		os.Exit(1)
	}
	port = os.Getenv("PORT")
	if port == "" {
		fmt.Printf("No port specified; defaulting to 8080")
		port = ":8080"
	}
	if port[0] != ':' {
		port = ":" + port
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

// WebPort returns the port which the web server should be running on, in the format :8080
func WebPort() string {
	return port
}
