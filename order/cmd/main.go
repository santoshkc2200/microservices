package main

import (
	"log"

	"gitHub.com/santoshkc2200/microservices/order/config"
	"gitHub.com/santoshkc2200/microservices/order/internal/adapters/db"
	"gitHub.com/santoshkc2200/microservices/order/internal/adapters/grpc"
	"gitHub.com/santoshkc2200/microservices/order/internal/adapters/payment"
	"gitHub.com/santoshkc2200/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	paymentAdapter, paymentError := payment.NewAdapter(config.GetPayme