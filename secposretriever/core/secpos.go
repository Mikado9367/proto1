package core

import (
	"context"
	"log"
	"time"

	pb "secposretriever/models/grpc/secpos"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// type MyServer struct {
// 	*config.Server
// }

type Server struct {
	pb.SecurityPosisionSettlementServiceServer
}

// retrive all from secpos_grpc.pb.go  SecurityPosisionSettlementServiceServer interface
func (s *Server) GetOneSecurityPosition(ctx context.Context, in *pb.SecPosRequest) (*pb.SecPosResponse, error) {

	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.
	// Coder l'accès à la base de donnée.

	log.Printf(("GetOneSecurityPosition invoked... %s\n"), in)

	secKey := &pb.SecurityPositionKey{
		Isin:            in.SearchCriteria.SecurityPositionKey.Isin,
		Account:         in.SearchCriteria.SecurityPositionKey.Account,
		RestrictionType: in.SearchCriteria.SecurityPositionKey.RestrictionType,
		ClientId:        in.SearchCriteria.SecurityPositionKey.ClientId,
	}

	secValue := &pb.SecurityPositionValue{
		PositionQuantity:    1234,
		PositionQuantitySod: 4,
		PeriodEvtReference:  "RTS",
		SettPositionTs:      &timestamppb.Timestamp{},
	}

	// une façon de faire Tradi.
	secValue.SettPositionTs.Seconds = time.Now().Unix()

	secPos := &pb.SecurityPosition{
		SecurityPositionKey:   secKey,
		SecurityPositionValue: secValue,
	}

	secResponse := pb.SecPosResponse{
		SecurityPosition: append([]*pb.SecurityPosition{}, secPos),
	}

	//	err := status.Error(codes.NotFound, "id was not found")
	//return nil, err

	return &secResponse, nil

	// return &secResponse, err
}
