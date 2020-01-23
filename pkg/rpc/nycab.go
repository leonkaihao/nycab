package rpc

import (
	"context"
	"time"

	pb "github.com/leonkaihao/nycab/api/proto/nycab"
)

// DoGetCabPickupCount implementation
func (r *Server) DoGetCabPickupCount(ctx context.Context, req *pb.GetCabPickupCountRequest) (*pb.GetCabPickupCountResponse, error) {
	if req == nil {
		return nil, errNullRequest
	}
	dayTime := time.Unix(0, req.DayTime)
	durDay := time.Hour * 24
	// range is the beginning of start date to the end of end date.
	dayTime = dayTime.Truncate(durDay).Add(durDay)
	ret, err := r.db.GetPickupCount(ctx, req.GetMedallions(), dayTime)
	if err != nil {
		return nil, err
	}
	return &pb.GetCabPickupCountResponse{
		DayTime: dayTime.UnixNano(),
		Info:    ret,
	}, err
}
