package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/leonkaihao/nycab/api/swagger/models"
	"github.com/leonkaihao/nycab/api/swagger/restapi/operations"
)

func (h *handler) GetPickupCount(params operations.GetCabsPickupsCountParams) middleware.Responder {
	result := []*models.CabPickupsCount{}
	for _, d := range params.Medallions {
		result = append(result, &models.CabPickupsCount{
			Medallions: d,
		})
	}
	return &operations.GetCabsPickupsCountOK{
		Payload: &models.GetCabsPickupsCountResponse{
			Code:   1,
			Result: result,
		},
	}
}
