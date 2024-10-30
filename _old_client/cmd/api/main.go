package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "client/secpos"
)

var addr string = "0.0.0.0:50051"

type Client struct {
	pb.SecurityPosisionSettlementServiceClient
}

func main() {

	//	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect : %v\n", err)
	}

	log.Printf(("start Settlement client service... %s\n"), addr)
	c := pb.NewSecurityPosisionSettlementServiceClient(conn)

	defer conn.Close()
	defer log.Printf(("closing... %s\n"), addr)

	log.Printf(("CallOne... %s\n"), addr)
	toto := CallOne(c)

	fmt.Printf("*************************  \n")
	fmt.Printf("toto :   %v  \n", toto)
	fmt.Printf("*************************  \n")
	fmt.Printf("nombre de position :   %v  \n", len(toto.SecurityPosition))
	fmt.Printf("*************************  \n")

	for _, s := range toto.SecurityPosition {
		fmt.Printf("*************************  \n")
		fmt.Printf("toto.timestamp :   %v  \n", s.SecurityPositionValue.SettPositionTs)

		// https://pkg.go.dev/google.golang.org/protobuf/types/known/timestamppb#Timestamp.AsTime
		ts := s.SecurityPositionValue.SettPositionTs.AsTime()

		formatted := ts.Format(time.RFC3339)

		fmt.Printf("*************************  \n")
		fmt.Printf("timestamp convertit :   %v  \n", formatted)
		fmt.Printf("*************************  \n")
	}

}
