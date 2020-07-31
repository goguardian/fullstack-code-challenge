package main

import (
	"log"

	"github.com/goguardian/Development/golang/ggg/jl"
	srsclient "github.com/goguardian/Development/golang/ggg/srs/client"

	"github.com/goguardian/Development/services/examples/go-grpc/config"
	gocoreapi "github.com/goguardian/Development/services/examples/go-grpc/pkg/datastore/go-core-api"
	"github.com/goguardian/Development/services/examples/go-grpc/service"
	"github.com/goguardian/Development/services/examples/go-grpc/transport/grpc"
	"github.com/goguardian/Development/services/examples/go-grpc/transport/http"
)

func main() {
	conf := config.Config

	var err error
	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Creating SRS client")
	srsClient, err := srsclient.New()
	if err != nil {
		return
	}

	logger.Info("Creating datastore client")
	datastoreClient, err := gocoreapi.New(&gocoreapi.Config{
		DatabaseAddress: conf.DatabaseAddress,
		ReadTimeout:     conf.GoCoreAPIReadTimeout,
	})
	if err != nil {
		return
	}

	log.Println("Creating service instance")
	server, err := service.New(&service.Config{
		DatastoreClient: datastoreClient,
	})
	if err != nil {
		return
	}

	go func() {
		log.Println("Creating HTTP listener")
		httpListener, err := http.New(&http.Config{
			GRPCListenPort:    conf.GRPCListenPort,
			HTTPListenAddress: conf.HTTPListenAddress,
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Fatal(httpListener.Listen())
	}()

	logger.WithFields(jl.Fields{
		"grpc_listen_address": conf.GRPCListenAddress,
	}).Info("Creating gRPC listener")
	grpcListener, err := grpc.New(&grpc.Config{
		GRPCListenAddress: conf.GRPCListenAddress,
		GRPCListenPort:    conf.GRPCListenPort,
		ServiceName:       conf.ServiceName,
		Server:            server,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Fatal(grpcListener.Listen())
}
