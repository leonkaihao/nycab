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
		found := true
		elem := &models.CabPickupsCount{
			Count:     &data.Count,
			Medallion: &data.Medallion,
			Found:     &found,
		}
		result = append(result, elem)
		delete(medallionsMap, data.GetMedallion())
	}
	// append not existed medallions
	for key := range medallionsMap {
		found := false
		elem := &models.CabPickupsCount{
			Medallion: &key,
			Found:     &found,
		}
		result = append(result, elem)
	}

	return &models.GetCabsPickupsCountResponse{
		Result: result,
	}
}
