package main

import (
	pb "client/secpos"
	"context"
	"fmt"
	"log"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/status"
)

func CallOne(c pb.SecurityPosisionSettlementServiceClient) (response *pb.SecPosResponse) {

	log.Printf("callOne invoker...\n")

	//crit := pb.SecPosRequest.SearchCriteria

	criteria := pb.SearchCriteria{
		BusinessPeriodType: pb.BusinessPeriodType_LAST,
		BusinessDate: &date.Date{
			Year:  2024,
			Month: 9,
			Day:   22,
		},
		SecurityPositionKey: &pb.SecurityPositionKey{
			Isin:            "ISIN1",
			Account:         "EXTACC1",
			RestrictionType: "RSTR1",
			ClientId:        "EXTCLI1"},
	}

	search := &pb.SecPosRequest{}
	search.SearchCriteria = &criteria

	res, err := c.GetOneSecurityPosition(context.Background(), search)

	st, ok := status.FromError(err)
	fmt.Printf("status code is : %v\n", st.Code())
	fmt.Printf("message is : %v\n", st.Message())

	if !ok {

		fmt.Printf("status code is : %v\n", st.Code())
		fmt.Printf("message is : %v\n", st.Message())

		// Error was not a status error
	}

	if err != nil {
		log.Fatalf("Failed to call GetOneSecurityPosition : %v\n", err)
	}

	return res
}
