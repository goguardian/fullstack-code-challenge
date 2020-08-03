package grpc

import (
	"errors"
	"fmt"
	"math"
	"net"
	gw "proto/fullstack-code-challenge"

	"google.golang.org/grpc"
)

// Config represents the configuration for a listener.
type Config struct {
	GRPCListenAddress string
	GRPCListenPort    uint
	Server            gw.GoGRPCServer
	ServiceName       string
}

// Listener describes the interface for interacting with a listener.
type Listener interface {
	Listen() error
}

// New creates and returns a new listener.
func New(conf *Config) (Listener, error) {
	if conf.GRPCListenAddress == "" {
		return nil, errors.New("invalid gRPC listen address")
	}
	if conf.GRPCListenPort == 0 {
		return nil, errors.New("invalid gRPC listen port")
	}
	if conf.Server == nil {
		return nil, errors.New("invalid server")
	}
	if conf.ServiceName == "" {
		return nil, errors.New("invalid service name")
	}

	return &listener{
		grpcListenAddress: conf.GRPCListenAddress,
		grpcListenPort:    conf.GRPCListenPort,
		serviceName:       conf.ServiceName,
		GoGRPCServer:      conf.Server,
	}, nil
}

type listener struct {
	gw.GoGRPCServer

	grpcListenAddress string
	grpcListenPort    uint
	serviceName       string
}

func (l *listener) Listen() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", l.grpcListenPort))
	if err != nil {
		return nil, err
	}

	// if err = registerService(l.grpcListenAddress, l.grpcListenPort); err != nil {
	// 	return nil, err
	// }

	grpcServerOptions := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(math.MaxUint32),
	}

	// grpcServe := srsserver.GetGRPCServer(grpcServerOptions...)
	grpcServe := grpc.NewServer(gprcOptions...)

	gw.RegisterGoGRPCServer(grpcServe, l)

	return grpcServe.Serve(listener)
}
