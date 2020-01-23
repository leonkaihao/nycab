package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/leonkaihao/nycab/api/swagger/restapi/operations"
	"github.com/leonkaihao/nycab/pkg/api"
	"github.com/leonkaihao/nycab/pkg/cache"
)

// Handler interface has all the handle functions and implemented in classified go files.
type Handler interface {
	DeleteCache(params operations.DeleteCabsPickupsCountCacheParams) middleware.Responder
	GetPickupCount(params operations.GetCabsPickupsCountParams) middleware.Responder
}

type handler struct {
	rpc *api.Client
	cch cache.Cache
}

// NewHandler ...
func NewHandler(client *api.Client, cch cache.Cache) Handler {
	return &handler{client, cch}
}
