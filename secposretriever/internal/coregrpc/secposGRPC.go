package coregrpc

import (
	"context"
	"log"

	"secposretriever/internal/coredb"
	"secposretriever/internal/tool"
	pb "secposretriever/models/grpc/secpos"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// retrive all from secpos_grpc.pb.go  SecurityPosisionSettlementServiceServer interface

type Server struct {
	pb.SecurityPosisionSettlementServiceServer
}

func (s *Server) ConvertGrpcToBusinessFormat(in *pb.SecPosRequest) (req coredb.RequestPosition) {

	req.Bic = in.SearchCriteria.SecurityPositionKey.ClientId
	req.Account = in.SearchCriteria.SecurityPositionKey.Account
	req.Isin = in.SearchCriteria.SecurityPositionKey.Isin
	req.Filter.Date = in.SearchCriteria.BusinessDate.String()
	req.Filter.Phase = in.SearchCriteria.BusinessPeriodType.String()

	return req
}

func (s *Server) GetOneSecurityPosition(ctx context.Context, in *pb.SecPosRequest) (resp *pb.SecPosResponse, err error) {

	log.Printf(("GetOneSecurityPosition invoked... %s\n"), in)

	myBusinessRequest := s.ConvertGrpcToBusinessFormat(in)

	if err = myBusinessRequest.CheckRequest(); err != nil {
		return &pb.SecPosResponse{}, err
	}

	if err := myBusinessRequest.Filter.IsFilterValid(); err != nil {
		return &pb.SecPosResponse{}, err
	}

	// retrieve data from Database
	myFeedBack, err := myBusinessRequest.GetPositionsFromDB()

	if err != nil {
		return &pb.SecPosResponse{}, err
	}

	lg := len(myFeedBack)
	if lg > 1 {
		err = coredb.ErrTooMuchRows
		return &pb.SecPosResponse{}, err
	}

	myResponse := &pb.SecPosResponse{}

	t, err := tool.StringIntoTime(myFeedBack[0].LastTimestamp)
	if err != nil {
		return &pb.SecPosResponse{}, err
	}

	// prepare result into gRPC format
	var oneSecurityPosition = &pb.SecurityPosition{
		SecurityPositionKey: &pb.SecurityPositionKey{
			Isin:            myFeedBack[0].Isin,
			Account:         myFeedBack[0].Account,
			RestrictionType: myFeedBack[0].Restrictiontype,
			ClientId:        myFeedBack[0].PartyBic,
		},
		SecurityPositionValue: &pb.SecurityPositionValue{
			PositionQuantity:    float64(myFeedBack[0].Quantity),
			PositionQuantitySod: 666,
			PeriodEvtReference:  myFeedBack[0].Phase,
			SettPositionTs:      timestamppb.New(t),
		},
	}

	myResponse.SecurityPosition = append(myResponse.SecurityPosition, oneSecurityPosition)

	return myResponse, nil
}
