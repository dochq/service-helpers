package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

type contextKey string

const (
	// Time  - Keepalive Server Parameter
	Time = 5 * time.Second
	// Timeout  - Keepalive Server Parameter
	Timeout = 5 * time.Second
	// MaxConnectionAgeGrace  - Keepalive Server Parameter
	MaxConnectionAgeGrace = 10 * time.Second
	// EmptyContextKey - empty context key
	EmptyContextKey = contextKey("")
)

// ServerKeepaliveParams - gRPC Server Keepalive Parameters
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

// ClientKeepaliveParams - gRPC Client Keepalive Parameters
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

// NewServer - create a gRPC Server
func NewServer(opt ...grpc.ServerOption) *grpc.Server {
	grpcServer := grpc.NewServer(
		append(opt, ServerKeepaliveParams)...,
	)
	return grpcServer
}

// Dial - establish a gRPC Connection
func Dial(target string, opt ...grpc.DialOption) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		"dns:///"+target,
		append(
			opt,
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
			ClientKeepaliveParams,
		)...,
	)
	return conn, err
}

// GetContextWithMetaAuth - add a key and a value into the Context Meta Data
func GetContextWithMetaAuth(ctx context.Context, key, value string) context.Context {
	return context.WithValue(
		metadata.NewOutgoingContext(
			ctx,
			metadata.New(map[string]string{
				key: value,
			}),
		),
		EmptyContextKey, "",
	)
}
