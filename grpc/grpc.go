package grpc

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	Time                  = 5 * time.Second
	Timeout               = 5 * time.Second
	MaxConnectionAgeGrace = 10 * time.Second
)

var ServerKeepaliveParams = grpc.KeepaliveParams(keepalive.ServerParameters{
	// After a duration of this time if the server doesn't see any activity it
	// pings the client to see if the transport is still alive.
	// If set below 1s, a minimum value of 1s will be used instead.
	Time: Time,
	// After having pinged for keepalive check, the server waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	Timeout: Timeout,
	// MaxConnectionAgeGrace is an additive period after MaxConnectionAge after
	// which the connection will be forcibly closed.
	MaxConnectionAgeGrace: MaxConnectionAgeGrace,
})

var ClientKeepaliveParams = grpc.WithKeepaliveParams(keepalive.ClientParameters{
	// After a duration of this time if the client doesn't see any activity it
	// pings the server to see if the transport is still alive.
	// If set below 10s, a minimum value of 10s will be used instead.
	Time: Time,
	// After having pinged for keepalive check, the client waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	Timeout: Timeout,
	// If true, client sends keepalive pings even with no active RPCs. If false,
	// when there are no active RPCs, Time and Timeout will be ignored and no
	// keepalive pings will be sent.
	PermitWithoutStream: true,
})

func NewServer(opt ...grpc.ServerOption) *grpc.Server {
	grpcServer := grpc.NewServer(
		ServerKeepaliveParams, opt,
	)
	return grpcServer
}
