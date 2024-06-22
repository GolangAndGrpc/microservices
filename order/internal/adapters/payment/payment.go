package payment

import (
	"context"
	"fmt"

	"github.com/GolangAndGrpc/microserviceApis/golang/payments"
	"github.com/GolangAndGrpc/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payments.PaymentClient
	conn *grpc.ClientConn
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error){
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8088", opts...)
	if err != nil {
		return nil, err
	}
	client := payments.NewPaymentClient(conn)
	return &Adapter{payment: client, conn: conn}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	fmt.Printf("Sending request to payment create")
	_, err := a.payment.Create(context.Background(), &payments.CreatePaymentRequest{
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