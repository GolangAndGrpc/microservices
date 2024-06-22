package main

import (
	"fmt"
	"log"

	"github.com/GolangAndGrpc/microservices/payment/config"
	"github.com/GolangAndGrpc/microservices/payment/internal/adaptars/db"
	grpcAdaptarpackate "github.com/GolangAndGrpc/microservices/payment/internal/adaptars/grpc"
	"github.com/GolangAndGrpc/microservices/payment/internal/application/core/api"
)

func main() {
	dbAdaptar, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		 log.Fatalf("Error while Creating dbAdaptar %v", dbAdaptar)
	}

	application := api.NewApplication(dbAdaptar)

	grpcAdaptar := grpcAdaptarpackate.NewAdapter(application, config.GetApplicationPort())
	fmt.Printf("Service running on port %v ..", config.GetApplicationPort())

	grpcAdaptar.Run()

}