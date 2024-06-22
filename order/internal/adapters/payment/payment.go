package payment

import (
	"context"
	"fmt"
	"time"

	"github.com/GolangAndGrpc/microserviceApis/golang/payments"
	"github.com/GolangAndGrpc/microservices/order/internal/application/core/domain"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payments.PaymentClient
	conn *grpc.ClientConn
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error){
	var opts []grpc.DialOption
	opts = append(opts, 
	grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
	)))

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8088", opts...)
	if err != nil {
		return nil, err
	}
	client := payments.NewPaymentClient(conn)
	return &Adapter{payment: client, conn: conn}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	fmt.Println("Sending request to payment service")
	_, err := a.payment.Create(ctx, &payments.CreatePaymentRequest{
		UserId: order.CustomerID,
		OrderId: order.ID,
		TotalPrice: order.TotalPrice(),
	})
	fmt.Printf("Error %v \n", err)
	return err
}

func (a *Adapter) Close() {
	a.conn.Close()
}