package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	gw "github.com/goguardian/fullstack-code-challenge/api/proto/gen/go/fullstack_code_challenge/v1"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// Config represents the configuration for a listener.
type Config struct {
	GRPCListenPort    uint
	HTTPListenAddress string
}

// Listener describes the interface for interacting with a listener.
type Listener interface {
	Listen() error
}

// New creates and returns a new listener.
func New(conf *Config) (Listener, error) {
	if conf.GRPCListenPort == 0 {
		return nil, errors.New("invalid gRPC listen port")
	}
	if conf.HTTPListenAddress == "" {
		return nil, errors.New("invalid HTTP listen address")
	}

	return &listener{
		grpcListenPort:    conf.GRPCListenPort,
		httpListenAddress: conf.HTTPListenAddress,
	}, nil
}

type listener struct {
	grpcListenPort    uint
	httpListenAddress string
}

func (l *listener) Listen() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	// Force inclusion of keys of structs with empty values.
	gmux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	mux.Handle("/", gmux)

	localGRPCAddress := fmt.Sprintf("127.0.0.1:%d", l.grpcListenPort)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterFullstackCodeChallengeHandlerFromEndpoint(ctx, gmux, localGRPCAddress, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(l.httpListenAddress, mux)
}
