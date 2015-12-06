package config

import (
	"fmt"
	"os"
)

var (
	env   string
	port  string
	dburl string
	mongoLoggingCluster string
)

// InitEnvs loads the environment variables we use to configure the rest of the app
func InitEnvs() {
	env = os.Getenv("ENV")
	if env != "dev" && env != "prod" {
		fmt.Printf("$ENV not set to either 'dev' or 'prod'")
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
	dburl = os.Getenv("DATABASE_URL")
	if dburl == "" {
		fmt.Printf("$DATABASE_URL not set")
		os.Exit(1)
	}
	mongoLoggingCluster = os.Getenv("MONGOLAB_URL")
	if mongoLoggingCluster == "" {
		fmt.Printf("$MONGOLAB_URL not set")
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

// WebPort returns the port which the web server should be running on, in the format :8080
func WebPort() string {
	return port
}

// PostgresURL returns the full url for connecting to postgres, including the username, password, and port
func PostgresURL() string {
	return dburl
}

// LoggingMongoURL returns the full url for connecting to the mongodb logging cluster
func LoggingMongoURL() string {
	return mongoLoggingCluster
}
