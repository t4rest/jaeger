package subprovider

import (
	"context"
	"fmt"
	"service1/proto/subscription"
	"time"

	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
)

// SubProvider represents an interface to the Subscription grpc service
type SubProvider interface {
	GetSubscription(ctx context.Context, subID string) (*subscription.SubscriptionResponse, error)
}

type subsProvider struct {
	conn      *grpc.ClientConn
	timeout   int
	subClient subscription.SubscriptionClient
}

// New subscription provider
func New() (*subsProvider, error) {

	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		return nil, fmt.Errorf("register error: %s", err)
	}

	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure(), grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		return nil, fmt.Errorf("grpc Dial error: %s", err)
	}

	return &subsProvider{
		conn:      conn,
		timeout:   10,
		subClient: subscription.NewSubscriptionClient(conn),
	}, nil
}

// GetSubscription .
func (p *subsProvider) GetSubscription(ctx context.Context, subID string) (*subscription.SubscriptionResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Duration(p.timeout)*time.Second)
	defer cancel()

	req := &subscription.SubscriptionRequest{SubscriptionId: subID}
	resp, err := p.subClient.GetSubscription(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to get subscription data: %s", err)
	}

	return resp, nil
}

// Close .
func (p *subsProvider) Close() {
	err := p.conn.Close()
	if err != nil {
		logrus.Error(err)
	}
}
