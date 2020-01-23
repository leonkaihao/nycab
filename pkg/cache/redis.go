package cache

import (
	"time"

	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/pkg/config"
)

type redis struct {
}

func (r *redis) WithData(medallion string, day time.Time, count int64) error {
	return nil
}

func (r *redis) WithMultiData(data []*pb.MedallionPickupInfo, day time.Time) error {
	return nil
}

func (r *redis) GetDataSingle(medallion string, day time.Time) (int64, error) {
	return 0, nil
}

func (r *redis) GetDataMulti(medallions []string, day time.Time) ([]*pb.MedallionPickupInfo, error) {
	return nil, nil
}

func (r *redis) Clear() error {
	return nil
}

// NewRedisCache ...
func NewRedisCache(conf *config.API) (Cache, error) {
	return &redis{}, nil
}
