package main

import (
	"log"
	"net"

	"broker/config"
	pb "broker/secpos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//import ("fmt")

// Settlement Broker API is the entry point for all Settlement services.

// Protocole is only with Grpc at this stage.
// Broker is Server regarding front and client with Retriever
type Server struct {
	pb.SecurityPosisionSettlementServiceServer
}

type Client struct {
	pb.SecurityPosisionSettlementServiceClient
}

func main() {

	// setup gRPC client
	var err error
	config.ConnGRPC, err = grpc.NewClient(config.Envs.SecPosRetrieverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	defer config.ConnGRPC.Close()
	defer log.Printf(("closing... %s\n"), config.Envs.SecPosRetrieverAddr)

	// setup gRPC server
	lis, err := net.Listen("tcp", config.Envs.AddrServerGRPC)

	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf(("starting Settlement broker service... %s\n"), config.Envs.AddrServerGRPC)

	s := grpc.NewServer()
	pb.RegisterSecurityPosisionSettlementServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {

		log.Fatalf("Failed to serve : %v\n", err)

	}

}
