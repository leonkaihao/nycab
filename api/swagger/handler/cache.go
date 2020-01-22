package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/leonkaihao/nycab/api/swagger/restapi/operations"
	log "github.com/sirupsen/logrus"
)

func (h *handler) DeleteCache(params operations.DeleteCabsPickupsCountCacheParams) middleware.Responder {
	log.Infof("%v, %v", params.StartDate.String(), params.EndDate.String())
	return &operations.DeleteCabsPickupsCountCacheOK{}
}
