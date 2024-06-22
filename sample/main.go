package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GolangAndGrpc/microserviceApis/golang/payments"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8088", opts...)
	if err != nil {
		log.Fatalf("Error : %v \n", err)
	}
	defer conn.Close()

	paymentClient := payments.NewPaymentClient(conn)
	res, err := paymentClient.Create(context.Background(), &payments.CreatePaymentRequest{
		UserId: 10,
		OrderId: 99,
		TotalPrice: 999,
	})
	if err != nil {
		fmt.Printf("Error is : %v \n", err)
	}
	fmt.Printf("Response is %v \n", res)
}
