package grpc_server

import (
	"net"
	"subscription/proto/subscription"

	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcSrv struct {
	grpcListener net.Listener
	grpcServer   *grpc.Server
}

// New return new grpc module
func New() *grpcSrv {
	return &grpcSrv{}
}

// Title .
func (srv *grpcSrv) Title() string {
	return "GRPC"
}

// Start .
func (srv *grpcSrv) Start() {
	var err error

	srv.grpcListener, err = net.Listen("tcp", ":8088")
	if err != nil {
		logrus.Fatalf("Error can't launch the grpc server on port: %s", ":8088")
	}

	// Register views to collect data.
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		logrus.Fatalf("register error: %s", err)
	}

	srv.grpcServer = grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	subscription.RegisterSubscriptionServer(srv.grpcServer, srv)

	err = srv.grpcServer.Serve(srv.grpcListener)
	if err != nil {
		logrus.Fatalf("grpcServer err: %s ", err)
	}
}

// Stop .
func (srv *grpcSrv) Stop() {
	srv.grpcServer.GracefulStop()
}
