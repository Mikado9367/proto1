package handlers

import (
	"context"
	"errors"
	"fmt"
	"front/config"
	pb "front/internal/clientgrpc/secpos"
	"front/internal/models/api"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/status"
)

// modifier pour retourner un apiAnswer qui correspond a une stuct hors GRPC. du coup faut aussi un convert gRPCtoAPI
func GetOneSecurityPosition(d api.ApiData) (response []api.SecurityPosition, err error) {

	search := ConvertApiToGrpc(d)

	if search == nil {

		err = errors.New(fmt.Sprintf("invalid ConvertApiToGrpc") + fmt.Sprintf("content is : %v\n", d))
		return response, err
	}

	c := pb.NewSecurityPosisionSettlementServiceClient(config.ConnGRPC)

	var feedback *pb.SecPosResponse

	feedback = &pb.SecPosResponse{}

	feedback, err = c.GetOneSecurityPosition(context.Background(), search)

	st, ok := status.FromError(err)

	if !ok {
		err = errors.New(fmt.Sprintf("status code is : %v\n", st.Code()) + fmt.Sprintf("message is : %v\n", st.Message()))
	}

	if err != nil {
		log.Fatalf("Failed to call GetOneSecurityPosition : %v\n", err)
	}

	if feedback == nil {
		log.Fatalf("Feedback nil from call GetOneSecurityPosition : %v\n", err)
		err = errors.New("Feedback nil from call GetOneSecurityPosition")

	}

	if ok && err != nil {
		response = ConvertGrpcToApi(feedback)
	}
	return response, err
}

func ConvertApiToGrpc(d api.ApiData) (r *pb.SecPosRequest) {

	t, _ := time.Parse("2006-01-02", d.SearchFilter.BusinessDate)

	r = &pb.SecPosRequest{}
	r.SearchCriteria = &pb.SearchCriteria{
		BusinessPeriodType: pb.BusinessPeriodType_LAST,
		BusinessDate: &date.Date{
			Year:  int32(t.Year()),
			Day:   int32(t.Day()),
			Month: int32(t.Month()),
		},
		SecurityPositionKey: &pb.SecurityPositionKey{
			Isin:            d.SecurityPositionKey.Isin,
			Account:         d.SecurityPositionKey.Account,
			RestrictionType: d.SecurityPositionKey.Restrictiontype,
			ClientId:        d.SecurityPositionKey.Clientbic,
		}}

	return r

}

func ConvertGrpcToApi(d *pb.SecPosResponse) (r []api.SecurityPosition) {

	// item.SecurityPositionKey.ClientId,

	for _, item := range d.SecurityPosition {

		fmt.Printf("toto %v", item)

		tmp := api.SecurityPosition{
			SecurityPositionKey: api.SecurityPositionKey{
				Isin:            item.SecurityPositionKey.Isin,
				Account:         item.SecurityPositionKey.Account,
				RestrictionType: item.SecurityPositionKey.RestrictionType,
				ClientId:        item.SecurityPositionKey.ClientId,
			},
			SecurityPositionValue: api.SecurityPositionValue{
				Position_quantity:     item.SecurityPositionValue.PositionQuantity,
				Position_quantity_sod: item.SecurityPositionValue.PositionQuantitySod,
				Period_evt_reference:  item.SecurityPositionValue.PeriodEvtReference,
				Sett_position_ts:      item.SecurityPositionValue.SettPositionTs.String(),
			},
		}

		r = append(r, tmp)

	}

	return r

}

// func CallOne(c pb.SecurityPosisionSettlementServiceClient) (response *pb.SecPosResponse) {

// 	log.Printf("callOne invoker...\n")

// 	//crit := pb.SecPosRequest.SearchCriteria

// 	criteria := pb.SearchCriteria{
// 		BusinessPeriodType: pb.BusinessPeriodType_LAST,
// 		BusinessDate: &date.Date{
// 			Year:  2024,
// 			Month: 9,
// 			Day:   22,
// 		},
// 		SecurityPositionKey: &pb.SecurityPositionKey{
// 			Isin:            "ISIN1",
// 			Account:         "EXTACC1",
// 			RestrictionType: "RSTR1",
// 			ClientId:        "EXTCLI1"},
// 	}

// 	search := &pb.SecPosRequest{}
// 	search.SearchCriteria = &criteria

// 	res, err := c.GetOneSecurityPosition(context.Background(), search)

// 	st, ok := status.FromError(err)
// 	fmt.Printf("status code is : %v\n", st.Code())
// 	fmt.Printf("message is : %v\n", st.Message())

// 	if !ok {

// 		fmt.Printf("status code is : %v\n", st.Code())
// 		fmt.Printf("message is : %v\n", st.Message())

// 		// Error was not a status error
// 	}

// 	if err != nil {
// 		log.Fatalf("Failed to call GetOneSecurityPosition : %v\n", err)
// 	}

// 	return res
// }
