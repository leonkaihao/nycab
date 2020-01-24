package cache

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/leonkaihao/nycab/api/proto/nycab"
)

type mem struct {
	countMapper *sync.Map
}

func (m *mem) WithData(medallion string, day time.Time, count int64) error {
	return m.withData(medallion, day, count)
}

func (m *mem) WithMultiData(info []*pb.MedallionPickupInfo, day time.Time) error {
	for _, data := range info {
		err := m.withData(data.GetMedallion(), day, data.GetCount())
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *mem) GetDataSingle(medallion string, day time.Time) (int64, error) {
	return m.getDataSingle(medallion, day)
}

func (m *mem) GetDataMulti(medallions []string, day time.Time) ([]*pb.MedallionPickupInfo, error) {
	ret := []*pb.MedallionPickupInfo{}
	for _, data := range medallions {
		count, err := m.getDataSingle(data, day)
		if err == nil {
			ret = append(ret, &pb.MedallionPickupInfo{
				Medallion: data,
				Count:     count,
			})
		}
	}
	return ret, nil
}

func (m *mem) Clear() error {
	m.countMapper = new(sync.Map)
	return nil
}

// NewMemCache ...
func NewMemCache() (Cache, error) {
	mem := &mem{
		countMapper: new(sync.Map),
	}
	return mem, nil
}

func (m *mem) withData(medallion string, day time.Time, count int64) error {
	key := formatMemCacheKey(medallion, day)
	m.countMapper.Store(key, count)
	return nil
}

func (m *mem) getDataSingle(medallion string, day time.Time) (int64, error) {
	key := formatMemCacheKey(medallion, day)
	if v, ok := m.countMapper.Load(key); ok {
		return v.(int64), nil
	}
	return 0, fmt.Errorf("getDataSingle: not found")
}

func formatMemCacheKey(medallion string, day time.Time) string {
	key := medallion
	dateFormat := "2006-01-02"
	key += ("@" + day.Format(dateFormat))
	return key
}
