package cache

import (
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/pkg/config"
	log "github.com/sirupsen/logrus"
)

type redis struct {
	pool *redigo.Pool
	conf *config.API
}

func (r *redis) WithData(medallion string, day time.Time, count int64) error {
	c, err := r.pool.Dial()
	if err != nil {
		return err
	}
	return r.withData(c, medallion, day, count)
}

func (r *redis) WithMultiData(info []*pb.MedallionPickupInfo, day time.Time) error {
	c, err := r.pool.Dial()
	if err != nil {
		return err
	}
	defer c.Close()
	for _, data := range info {
		err := r.withData(c, data.GetMedallion(), day, data.GetCount())
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *redis) GetDataSingle(medallion string, day time.Time) (int64, error) {
	c, err := r.pool.Dial()
	if err != nil {
		return 0, err
	}
	defer c.Close()
	return r.getDataSingle(c, medallion, day)
}

func (r *redis) GetDataMulti(medallions []string, day time.Time) ([]*pb.MedallionPickupInfo, error) {
	c, err := r.pool.Dial()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	ret := []*pb.MedallionPickupInfo{}
	for _, data := range medallions {
		count, err := r.getDataSingle(c, data, day)
		if err == nil {
			ret = append(ret, &pb.MedallionPickupInfo{
				Medallion: data,
				Count:     count,
			})
		}
	}
	return ret, nil
}

func (r *redis) Clear() error {
	c, err := r.pool.Dial()
	if err != nil {
		return err
	}
	defer c.Close()
	_, err = c.Do("FLUSHDB")
	if err != nil {
		return fmt.Errorf("Clear failed, %v", err)
	}
	return nil
}

// NewRedisCache ...
func NewRedisCache(conf *config.API) (Cache, error) {
	if conf == nil {
		return nil, fmt.Errorf("redis missing config param")
	}

	url := conf.RedisURL

	pool := &redigo.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redigo.Conn, error) {
			conn, err := redigo.Dial("tcp", url)
			if err != nil {
				log.Errorf("ERROR: fail init redis: %s", err.Error())
			}
			return conn, err
		},
	}
	c, err := pool.Dial()
	if err != nil {
		return nil, fmt.Errorf("failed connect to redis")
	}
	defer c.Close()

	return &redis{pool, conf}, nil
}

func formatRedisKey(medallion string, day time.Time) string {
	key := medallion
	dateFormat := "2006-01-02"
	key += ("@" + day.Format(dateFormat))
	return key
}

func (r *redis) withData(c redigo.Conn, medallion string, day time.Time, count int64) error {
	key := formatRedisKey(medallion, day)
	_, err := c.Do("SET", key, count)
	if err != nil {
		return fmt.Errorf("WithData failed, %v", err)
	}
	return nil
}

func (r *redis) getDataSingle(c redigo.Conn, medallion string, day time.Time) (int64, error) {
	key := formatRedisKey(medallion, day)
	ret, err := redigo.Int64(c.Do("GET", key))
	if err != nil {
		return 0, fmt.Errorf("getDataSingle failed, %v", err)
	}
	return ret, nil
}
