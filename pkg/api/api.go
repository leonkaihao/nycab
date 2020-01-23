package api

import (
	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/pkg/config"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Client serving for rpc api
type Client struct {
	NYCab pb.NYCabClient
}

// NewNYCabClient get access to service.nycab
func NewNYCabClient(conf *config.API) (pb.NYCabClient, error) {
	nyCabURL := conf.SvcURL
	// Set up a connection to the server.
	log.Infoln("NYCab grpc host:" + nyCabURL)
	conn, err := grpc.Dial(nyCabURL,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	return pb.NewNYCabClient(conn), nil
}

// NewClient wrap all the clients will be visited.
func NewClient(conf *config.API) (*Client, error) {
	nyCab, err := NewNYCabClient(conf)
	if err != nil {
		return nil, err
	}
	return &Client{
		NYCab: nyCab,
	}, nil
}
