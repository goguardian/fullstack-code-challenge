package main

import (
	"log"

	"github.com/goguardian/fullstack-code-challenge/api/config"
	"github.com/goguardian/fullstack-code-challenge/api/pkg/datastore/mysql"
	"github.com/goguardian/fullstack-code-challenge/api/service"
	"github.com/goguardian/fullstack-code-challenge/api/transport/grpc"
	"github.com/goguardian/fullstack-code-challenge/api/transport/http"
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
	if err != nil {
		return
	}

	logger.Info("Creating datastore client")
	datastoreClient, err := mysql.New(&mysql.Config{
		DatabaseAddress: conf.DatabaseAddress,
		ReadTimeout:     conf.DatabaseReadTimeout,
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

	log.Println("Creating gRPC listener")
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
