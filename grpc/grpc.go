package grpc

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	Time                  = 5 * time.Second // wait X seconds, then send ping if there is no activity
	Timeout               = 5 * time.Second // wait for ping back
	MaxConnectionAgeGrace = 10 * time.Second
)

var keepaliveParams = grpc.KeepaliveParams(keepalive.ServerParameters{
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
