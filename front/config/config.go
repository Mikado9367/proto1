package config

import (
	"os"

	"google.golang.org/grpc"
)

const (
	WildCard = "-"
)

// config grpc
var AddrServerGRPC string = "0.0.0.0:50051"
var ConnGRPC *grpc.ClientConn

// config http
type Config struct {
	APIrelease string
	PublicHost string
	Port       string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		APIrelease: getEnv("RELEASE", "/api/settlement/v1"),
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", ":8080"),
	}
}

// const (
// 	twoDaysInSeconds = 60 * 60 * 24 * 2
// )

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
