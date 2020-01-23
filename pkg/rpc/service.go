package rpc

import (
	"context"
	"net"

	pb "github.com/leonkaihao/nycab/api/proto/nycab"
	"github.com/leonkaihao/nycab/pkg/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Server implement proto NYCabServer
type Server struct {
	health *health.Server
	db     db.Service
}

// NewPRCServer implementation
func NewPRCServer(db db.Service) pb.NYCabServer {
	svr := &Server{
		health: health.NewServer(),
		db:     db,
	}
	return svr
}

// Start RPC listening service at given network listener
// Terminates when ctx.Done is received
func Start(ctx context.Context, t *Server, lis net.Listener) {
	s := grpc.NewServer()

	pb.RegisterNYCabServer(s, t)
	grpc_health_v1.RegisterHealthServer(s, t.health)
	t.health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()
	log.Infof("RPC service is listening at %v...\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
