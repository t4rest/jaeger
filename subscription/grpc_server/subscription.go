package grpc_server

import (
	"context"
	"subscription/proto/subscription"
	"time"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetSubscription .
func (srv *grpcSrv) GetSubscription(ctx context.Context, sub *subscription.SubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	logrus.Info("GetSubscription")

	if sub == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	err := sub.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	time.Sleep(1 * time.Second)

	resp := &subscription.SubscriptionResponse{
		UserId:           "UserId",
		SubscriptionId:   "SubscriptionId",
		ConnectionId:     "ConnectionId",
		SubscriptionType: "SubscriptionType",
		ClientState:      "ClientState",
	}

	return resp, nil
}

// GetUserGroup .
func (srv *grpcSrv) GetUserGroup(ctx context.Context, sub *subscription.GroupRequest) (*subscription.GroupResponse, error) {
	if sub == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	err := sub.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	resp := &subscription.GroupResponse{
		UserId: sub.UserId,
	}

	return resp, nil
}
