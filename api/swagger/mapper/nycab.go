package mapper

import (
	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/api/swagger/models"
)

// MapPickupCountResponseToJSON ...
func MapPickupCountResponseToJSON(medallions []string, resp *pb.GetCabPickupCountResponse) *models.GetCabsPickupsCountResponse {
	if resp == nil {
		return nil
	}
	// create a medallions map to ensure
	// those medallions which does not exist in response
	medallionsMap := map[string]bool{}
	for _, m := range medallions {
		medallionsMap[m] = true
	}

	result := []*models.CabPickupsCount{}
	for _, data := range resp.Info {
		elem := &models.CabPickupsCount{
			Count:     data.GetCount(),
			Medallion: data.GetMedallion(),
			Found:     true,
		}
		result = append(result, elem)
		delete(medallionsMap, data.GetMedallion())
	}
	// append not existed medallions
	for key := range medallionsMap {
		elem := &models.CabPickupsCount{
			Medallion: key,
			Found:     false,
		}
		result = append(result, elem)
	}

	return &models.GetCabsPickupsCountResponse{
		Code:   0,
		Result: result,
	}
}
