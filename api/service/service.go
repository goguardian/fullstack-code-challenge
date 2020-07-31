package service

import (
	"errors"
	gw "proto/services/gateways/examples-go-grpc"

	"github.com/goguardian/Development/services/examples/go-grpc/pkg/datastore"
)

// Config represents the configuration for a service instance.
type Config struct {
	DatastoreClient datastore.Client
}

// New creates and returns a new service instance.
func New(conf *Config) (gw.GoGRPCServer, error) {
	if conf.DatastoreClient == nil {
		return nil, errors.New("invalid datastore client")
	}

	return &service{
		datastoreClient: conf.DatastoreClient,
	}, nil
}

type service struct {
	datastoreClient datastore.Client
}
