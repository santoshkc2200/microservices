package main

import (
	"log"

	"github.com/santoshkc2200/microservices/payment/config"
	"github.com/santoshkc2200/microservices/payment/internal/adapters/db"
	"github.com/santoshkc2200/microservices/payment/internal/adapters/grpc"
	"github.com/santoshkc2200/microservices/payment/internal/application/core/api"
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
