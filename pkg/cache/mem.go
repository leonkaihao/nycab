package cache

import (
	"fmt"
	"time"

	pb "github.com/leonkaihao/nycab/api/proto/nycab"
)

type mem struct {
}

func (m *mem) WithData(medallion string, day time.Time, count int64) error {
	return nil
}

func (m *mem) WithMultiData(data []*pb.MedallionPickupInfo, day time.Time) error {
	return nil
}

func (m *mem) GetDataSingle(medallion string, day time.Time) (int64, error) {
	return 0, nil
}

func (m *mem) GetDataMulti(medallions []string, day time.Time) ([]*pb.MedallionPickupInfo, error) {
	return nil, nil
}

func (m *mem) Clear() error {
	return nil
}

// NewMemCache ...
func NewMemCache() (Cache, error) {
	return nil, fmt.Errorf("not implement")
}
