package handler

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/api/swagger/mapper"
	"github.com/leonkaihao/nycab/api/swagger/restapi/operations"
	log "github.com/sirupsen/logrus"
)

func (h *handler) GetPickupCount(params operations.GetCabsPickupsCountParams) middleware.Responder {
	if h.rpc == nil {
		return PreConditionFailedError("cannot connect to rpc service")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	date, medallions, refresh := params.Date, params.Medallions, params.Refresh
	log.Infof("GetPickupCount Date:%v Medallions:%v Refresh:%v", date.String(), medallions, refresh)
	var (
		infoArr  []*pb.MedallionPickupInfo
		notFound []string
		err      error
	)
	if h.cch == nil {
		log.Errorln("cannot connect to cache service")
	}
	if h.cch != nil && (refresh == nil || *refresh == false) {
		infoArr, notFound, err = h.getPickupCountFromCache(ctx, medallions, time.Time(date))
		if err != nil {
			return InternalServerError(err.Error())
		}
		if len(notFound) != 0 { // if some medallions were not found in cache, search from DB
			infoArrRPC, err := h.getPickupCountFromRPC(ctx, notFound, time.Time(date))
			if err != nil {
				return InternalServerError(err.Error())
			}
			infoArr = append(infoArr, infoArrRPC...)
		}
	} else {
		infoArr, err = h.getPickupCountFromRPC(ctx, medallions, time.Time(date))
		if err != nil {
			return InternalServerError(err.Error())
		}
	}
	dateInt := time.Time(date).UnixNano()
	respJSON := mapper.MapPickupCountResponseToJSON(medallions, &pb.GetCabPickupCountResponse{
		DayTime: dateInt,
		Info:    infoArr,
	})
	return &operations.GetCabsPickupsCountOK{
		Payload: respJSON,
	}
}

// getPickupCountFromCache search cache, return found and not found item
func (h *handler) getPickupCountFromCache(ctx context.Context, medallions []string, date time.Time) (info []*pb.MedallionPickupInfo, medallionsNotFound []string, err error) {
	info, err = h.cch.GetDataMulti(medallions, date)
	if err != nil {
		return nil, medallions, err
	}
	medallionsMap := map[string]bool{}
	for _, medallion := range medallions {
		medallionsMap[medallion] = true
	}
	for _, elem := range info {
		delete(medallionsMap, elem.Medallion)
	}
	for key := range medallionsMap {
		medallionsNotFound = append(medallionsNotFound, key)
	}
	return info, medallionsNotFound, nil
}

// getPickupCountFromRPC search RPC, return found item and update cache
func (h *handler) getPickupCountFromRPC(ctx context.Context, medallions []string, date time.Time) ([]*pb.MedallionPickupInfo, error) {
	resp, err := h.rpc.NYCab.DoGetCabPickupCount(ctx, &pb.GetCabPickupCountRequest{
		DayTime:    date.UnixNano(),
		Medallions: medallions})
	if err != nil {
		return nil, err
	}
	// update cache concurrently
	go func() {
		if h.cch == nil {
			return
		}
		err = h.cch.WithMultiData(resp.GetInfo(), date)
		if err != nil {
			log.Error(err)
		}
	}()
	return resp.GetInfo(), nil
}
