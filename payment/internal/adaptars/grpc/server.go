package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/GolangAndGrpc/microserviceApis/golang/payments"
	"github.com/GolangAndGrpc/microservices/payment/config"
	"github.com/GolangAndGrpc/microservices/payment/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct{
	api ports.ApiPort
	port int
	payments.UnimplementedPaymentServer
}

func NewAdapter(api ports.ApiPort, port int) *Adapter{
	return &Adapter{api:api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp",fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	payments.RegisterPaymentServer(grpcServer,a)
	if config.GetEnv() == "dev" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("failed to serve grpc on port ")
	}

}
