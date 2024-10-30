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
	//	MyAddr              string
	AddrServerGRPC      string
	ConnGRPC            *grpc.ClientConn
	SecPosRetrieverAddr string
	// CashBalRetrieverAddr string
}

// essentials
var Srv *grpc.Server
var Envs = initConfig()
var ConnGRPC *grpc.ClientConn

func initConfig() Config {
	return Config{
		AddrServerGRPC:      getEnv("SettlementBrokerServer", "0.0.0.0:50051"),
		SecPosRetrieverAddr: getEnv("SettlementBrokerClient", "0.0.0.0:50052"),
		//		MyAddr:         getEnv("MyAddr", "0.0.0.0:50051"),
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
