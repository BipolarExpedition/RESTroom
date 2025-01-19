package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	//log "github.com/sirupsen/logrus"
)

// TODO: Setup dynamic versioning
//
// var (
// 	Version = "dev"
// 	Commit  = "unknown"
// 	BuildDate = "unknown"
// )

// TODO: Must configure Makefile to build the binary with the version information

// TODO: Setup logger
// func setup_logger(log_level string) {

// }

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value

}

type EnvSettings struct {
	LogLevel    string
	StderrLevel string
	IP          string
	Port        string
	Debug       bool
	AuthFile    string
}

func GetEnvSettings() *EnvSettings {
	// Get variables from the environment
	// RESTROOM_LOG_LEVEL, RESTROOM_PORT, RESTROOM_IP, RESTROOM_STDERR_LEVEL
	// RESTROOM_DEBUG, RESTROOM_AUTH_FILE

	// Get log level and stderr level
	log_level := strings.ToLower(GetEnvOrDefault("RESTROOM_LOG_LEVEL", "warn"))
	stderr_level := strings.ToLower(GetEnvOrDefault("RESTROOM_STDERR_LEVEL", "error"))
	// setup_logger(log_level)

	// Get address and port to bind to
	restroom_port := GetEnvOrDefault("RESTROOM_PORT", "9988")
	restroom_ip := os.Getenv("RESTROOM_IP")

	// Get if debug mode is enabled
	debug_mode :=
		GetEnvOrDefault("RESTROOM_DEBUG", "false")

	debug_bool := false
	switch strings.ToLower(strings.TrimSpace(debug_mode)) {
	case "true", "t", "on", "enabled", "1", "yes", "y":
		debug_bool = true
	default:
		debug_bool = false
	}

	// Get the auth file
	auth_file := GetEnvOrDefault("RESTROOM_AUTH_FILE", "/etc/restroom/auth.yaml")

	return &EnvSettings{
		LogLevel:    log_level,
		StderrLevel: stderr_level,
		IP:          restroom_ip,
		Port:        restroom_port,
		Debug:       debug_bool,
		AuthFile:    auth_file,
	}
}

func main() {

	envSettings := GetEnvSettings()

	restroom_address := envSettings.IP
	if len(envSettings.Port) > 0 {
		restroom_address = fmt.Sprintf("%s:%s", envSettings.IP, envSettings.Port)
	}

	// TODO: Setup Echo server and available services
	APIServer := echo.New()
	APIServer.Server.Addr = restroom_address

	// TODO: Start the server

	fmt.Println("Restroom")
}
