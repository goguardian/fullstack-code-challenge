package service

import (
	"errors"

	gw "github.com/goguardian/fullstack-code-challenge/api/proto/gen/go/fullstack_code_challenge/v1"

	"github.com/goguardian/fullstack-code-challenge/api/pkg/datastore"
)

// Config represents the configuration for a service instance.
type Config struct {
	DatastoreClient datastore.Client
}

// New creates and returns a new service instance.
func New(conf *Config) (gw.FullstackCodeChallengeServer, error) {
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
