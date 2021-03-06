// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteCabsPickupsCountCacheOKCode is the HTTP code returned for type DeleteCabsPickupsCountCacheOK
const DeleteCabsPickupsCountCacheOKCode int = 200

/*DeleteCabsPickupsCountCacheOK OK

swagger:response deleteCabsPickupsCountCacheOK
*/
type DeleteCabsPickupsCountCacheOK struct {
}

// NewDeleteCabsPickupsCountCacheOK creates DeleteCabsPickupsCountCacheOK with default headers values
func NewDeleteCabsPickupsCountCacheOK() *DeleteCabsPickupsCountCacheOK {

	return &DeleteCabsPickupsCountCacheOK{}
}

// WriteResponse to the client
func (o *DeleteCabsPickupsCountCacheOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
