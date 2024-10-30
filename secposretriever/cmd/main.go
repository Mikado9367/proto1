package main

import (
	"log"
	"net"

	"secposretriever/core"

	"secposretriever/config"
	pb "secposretriever/models/grpc/secpos"

	"google.golang.org/grpc"
)

//import ("fmt")

// Settlement security position retriever.
// Protocole is on Grpc at this stage
// Database is Postgres

// var addr string = "0.0.0.0:50052"

var MyDB config.DBstruct

func main() {

	// DB connect

	lis, err := net.Listen("tcp", config.Envs.AddrServerGRPC)

	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}

	log.Printf(("starting Settlement security position retriever service... %s\n"), config.Envs.AddrServerGRPC)

	Srv := grpc.NewServer()
	// pb.RegisterSecurityPosisionSettlementServiceServer(Srv, &config.Server{})
	pb.RegisterSecurityPosisionSettlementServiceServer(Srv, &core.Server{})

	// Server listening gRPC request
	if err = Srv.Serve(lis); err != nil {

		log.Fatalf("Failed to serve : %v\n", err)

	}

}
