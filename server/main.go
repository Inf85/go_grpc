package main

import (
	"github.com/joho/godotenv"
	"grpc/helpers"
	"grpc/server/serverFactory"
)

func init() {
	_ = godotenv.Load()
}
func main() {
	serverType := helpers.GetEnvDefault("SERVER_TYPE", "")
	if serverType == "http" {
		serverFactory.GetServerFactoryBuilder("http").Build()
	}
}
