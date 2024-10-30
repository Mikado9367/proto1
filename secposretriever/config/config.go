package config

import (
	"os"

	"google.golang.org/grpc"
)

const (
	WildCard = "-"
)

// config grpc Settlement security retriever
// var AddrServerGRPC string = "0.0.0.0:50052"
// var ConnGRPC *grpc.ClientConn

// config grpc
type Config struct {
	AddrServerGRPC string
	//	ConnGRPC       *grpc.ClientConn
	DBCONN string
	// DBUser     string
	// DBPassword string
	// DBAddress  string
	DBName   string
	WILDCARD string
}

// essentials
var Srv *grpc.Server
var Envs = initConfig()

func initConfig() Config {
	return Config{
		AddrServerGRPC: getEnv("AddrServerGRPC", "0.0.0.0:50052"),
		DBCONN:         getEnv("DBCONN", "postgres://postgres:postgres@172.21.0.3:5432/t2sSettInfoDB?"),
		WILDCARD:       getEnv("WILDCARD", "-"),
		// DBUser:     getEnv("DB_USER", "postgres"),
		// DBPassword: getEnv("DB_PASSWORD", "postgres"),
		// DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "172.0.0.3"), getEnv("DB_PORT", "5432")),
		DBName: getEnv("DB_NAME", "t2sSettInfoDB"),
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
