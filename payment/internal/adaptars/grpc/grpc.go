package grpc

import (
	"context"
	"fmt"

	"github.com/GolangAndGrpc/microserviceApis/golang/payments"
	"github.com/GolangAndGrpc/microservices/payment/internal/application/core/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payments.CreatePaymentRequest) (*payments.CreatePaymentResponse, error) {
	log.WithContext(ctx).Info("Creating Payment........")
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx,newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failded to charge %v", err)).Err()
	}

	return &payments.CreatePaymentResponse{PaymentId: result.ID}, nil
}