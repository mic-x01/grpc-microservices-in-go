package main

import (
	"log"

	"github.com/mic-x01/grpc-microservices-in-go/order/config"
	"github.com/mic-x01/grpc-microservices-in-go/order/internal/adapters/db"
	"github.com/mic-x01/grpc-microservices-in-go/order/internal/adapters/grpc"
	"github.com/mic-x01/grpc-microservices-in-go/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
