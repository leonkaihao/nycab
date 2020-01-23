package main

import (
	"context"
	"github.com/leonkaihao/nycab/pkg/rpc"
	"net"
	"sync"

	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/pkg/config"
	"github.com/leonkaihao/nycab/pkg/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port        = "0.0.0.0:8081"
	serviceName = "service.nycab"
)

func getConfig() (conf *config.Service, err error) {
	v, err := config.LoadConfig("service.nycab")
	if err != nil {
		return
	}
	conf, err = config.GetServiceConfig(v)
	if err != nil {
		return
	}
	return
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := getConfig()
	if err != nil {
		log.Fatalln("Cannot read config: ", err)
	}

	log.Infof("service config: %+v", conf)

	db, err := db.NewService(conf)
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()

		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		nyCabServer := rpc.NewPRCServer(db)
		pb.RegisterNYCabServer(s, nyCabServer)
		reflection.Register(s)

		go func() {
			<-ctx.Done()
			s.GracefulStop()
		}()
		log.Infof("RPC service is listening at %v...\n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Wait()
	log.Warning("Server has exited.")
}
