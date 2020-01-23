package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/pkg/config"
	log "github.com/sirupsen/logrus"
	"time"
)

// Service ...
type Service interface {
	GetPickupCount(ctx context.Context, medallions []string, tm time.Time) ([]*pb.MedallionPickupInfo, error)
}

type service struct {
	db *sqlx.DB
}

func (svc *service) GetPickupCount(ctx context.Context, medallions []string, tm time.Time) ([]*pb.MedallionPickupInfo, error) {

	if len(medallions) == 0 {
		return nil, errEmptyMedallions
	}
	dateTimeFormat := "2006-01-02 15:04:05"
	tmStart, tmEnd := tm.Format(dateTimeFormat), tm.Add(time.Hour*24).Format(dateTimeFormat)

	infoArr := []*pb.MedallionPickupInfo{}
	schemaCount := "select medallion, COUNT(medallion) as count from cab_trip_data where medallion in ("

	schemaCount += ("'" + medallions[0] + "'")
	for i := 1; i < len(medallions); i++ {
		schemaCount += (",'" + medallions[i] + "'")
	}
	schemaCount += ") and pickup_datetime>='" + tmStart + "' and pickup_datetime<'" + tmEnd
	schemaCount += "' group by medallion;"

	log.Infof("GetPickupCount: medallions %+v from %v", medallions, tmStart)
	err := svc.db.Select(&infoArr, schemaCount)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return infoArr, nil
}

// NewService ...
func NewService(conf *config.Service) (Service, error) {
	if conf == nil {
		return nil, errArgsLost
	}
	fullURL := conf.UserName + ":" + conf.Password + "@(" + conf.DBHost + ":" + conf.DBPort + ")/nycab"
	log.Infof("db url string: %v", fullURL)
	db, err := sqlx.Connect("mysql", fullURL)
	if err != nil {
		log.Fatalln(err)
	}

	svc := &service{db}
	return svc, err
}
