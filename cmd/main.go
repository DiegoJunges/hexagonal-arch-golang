package main

import (
	"hexagonal-arch-go/internal/adapters/app/api"
	"hexagonal-arch-go/internal/adapters/core/arithmetic"
	"hexagonal-arch-go/internal/adapters/framework/right/db"
	"hexagonal-arch-go/internal/ports"
	"log"
	"os"

	gRPC "hexagonal-arch-go/internal/adapters/framework/left/grpc"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatal("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
