// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteCabsPickupsCountCacheHandlerFunc turns a function with the right signature into a delete cabs pickups count cache handler
type DeleteCabsPickupsCountCacheHandlerFunc func(DeleteCabsPickupsCountCacheParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCabsPickupsCountCacheHandlerFunc) Handle(params DeleteCabsPickupsCountCacheParams) middleware.Responder {
	return fn(params)
}

// DeleteCabsPickupsCountCacheHandler interface for that can handle valid delete cabs pickups count cache params
type DeleteCabsPickupsCountCacheHandler interface {
	Handle(DeleteCabsPickupsCountCacheParams) middleware.Responder
}

// NewDeleteCabsPickupsCountCache creates a new http.Handler for the delete cabs pickups count cache operation
func NewDeleteCabsPickupsCountCache(ctx *middleware.Context, handler DeleteCabsPickupsCountCacheHandler) *DeleteCabsPickupsCountCache {
	return &DeleteCabsPickupsCountCache{Context: ctx, Handler: handler}
}

/*DeleteCabsPickupsCountCache swagger:route DELETE /cabs/pickups/count/cache deleteCabsPickupsCountCache

Method to clear the cache

*/
type DeleteCabsPickupsCountCache struct {
	Context *middleware.Context
	Handler DeleteCabsPickupsCountCacheHandler
}

func (o *DeleteCabsPickupsCountCache) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCabsPickupsCountCacheParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
