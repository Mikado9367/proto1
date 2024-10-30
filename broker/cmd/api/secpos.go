package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"broker/config"
	pb "broker/secpos"

	"google.golang.org/grpc/status"
)

// retrive all from secpos_grpc.pb.go  SecurityPosisionSettlementServiceServer interface

func (s *Server) GetOneSecurityPosition(ctx context.Context, in *pb.SecPosRequest) (feedback *pb.SecPosResponse, err error) {

	feedback, err = s.ForwardAndSendBack(ctx, in)
	return feedback, err

}

func (s *Server) ForwardAndSendBack(ctx context.Context, in *pb.SecPosRequest) (feedback *pb.SecPosResponse, err error) {

	log.Printf(("GetOneSecurityPosition invoked... %s\n"), in)

	c := pb.NewSecurityPosisionSettlementServiceClient(config.ConnGRPC)
	feedback, err = c.GetOneSecurityPosition(context.Background(), in)

	st, ok := status.FromError(err)

	if !ok {
		err = errors.New(fmt.Sprintf("status code is : %v\n", st.Code()) + fmt.Sprintf("message is : %v\n", st.Message()))
	}

	if err != nil {
		errors.Join(err, errors.New(fmt.Sprintf("status code is : %v\n", st.Code())+fmt.Sprintf("message is : %v\n", st.Message())))
		//		log.Fatalf("Failed to call GetOneSecurityPosition : %v\n", err)
	}

	return feedback, err

}
