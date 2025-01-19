package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value

}

func main() {
	// TODO: Get variables from the environment
	GetEnvOrDefault("RESTROOM_PORT", "9988")

	// TODO: Setup logger

	// TODO: Setup Echo server and available services
	APIServer := echo.New()
	APIServer.Server.Addr = ":9988"

	// TODO: Start the server

	fmt.Println("Restroom")
}
