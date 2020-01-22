package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/leonkaihao/nycab/api/swagger/restapi/operations"
)

// Handler interface has all the handle functions and implemented in classified go files.
type Handler interface {
	DeleteCache(params operations.DeleteCabsPickupsCountCacheParams) middleware.Responder
	GetPickupCount(params operations.GetCabsPickupsCountParams) middleware.Responder
}

type handler struct {
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{}
}
