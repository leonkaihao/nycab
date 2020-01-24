package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/leonkaihao/nycab/api/swagger/restapi/operations"
)

func (h *handler) DeleteCache(params operations.DeleteCabsPickupsCountCacheParams) middleware.Responder {
	if h.cch != nil {
		err := h.cch.Clear()
		if err != nil {
			return InternalServerError("failed to clear cache," + err.Error())
		}
	}
	return &operations.DeleteCabsPickupsCountCacheOK{}
}
