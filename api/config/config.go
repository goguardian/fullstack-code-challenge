package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config represents the runtime configuration for the service.
var Config *configuration

type configuration struct {
	DatabaseAddress     string
	DatabaseReadTimeout time.Duration

	GRPCListenPort    uint
	GRPCListenAddress string

	HTTPListenHost    string
	HTTPListenPort    int
	HTTPListenAddress string

	ServiceName string
}

func init() {
	Config = &configuration{
		DatabaseAddress:     getEnvOrDefault("DATABASE_ADDRESS", ""),
		DatabaseReadTimeout: time.Duration(getEnvIntOrDefault("DATABASE_READ_TIMEOUT", 10)) * time.Second,

		GRPCListenAddress: getEnvOrDefault("GRPC_LISTEN_ADDRESS", "grpc://fullstack-code-challenge"),
		GRPCListenPort:    uint(getEnvIntOrDefault("GRPC_LISTEN_PORT", 8889)),

		HTTPListenHost: getEnvOrDefault("HTTP_LISTEN_HOST", "0.0.0.0"),
		HTTPListenPort: getEnvIntOrDefault("HTTP_LISTEN_PORT", 8890),

		ServiceName: "fullstack-code-challenge",
	}

	Config.HTTPListenAddress = fmt.Sprintf("%s:%d", Config.HTTPListenHost, Config.HTTPListenPort)
}

// getEnvOrDefault returns an environment variable value if found, otherwise
// it returns the provided default.
func getEnvOrDefault(variable string, defaultValue string) string {
	if val := os.Getenv(variable); val != "" {
		return val
	}

	return defaultValue
}

func getEnvIntOrDefault(variable string, defaultValue int) int {
	if val := os.Getenv(variable); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}

	return defaultValue
}
