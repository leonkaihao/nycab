package cache

import (
	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"time"
)

// Cache defines the interface for all cache implementations.
type Cache interface {
	WithData(medallion string, day time.Time, count int64) error
	WithMultiData(data []*pb.MedallionPickupInfo, day time.Time) error
	GetDataSingle(medallion string, day time.Time) (int64, error)
	GetDataMulti(medallions []string, day time.Time) ([]*pb.MedallionPickupInfo, error)
	Clear() error
}
