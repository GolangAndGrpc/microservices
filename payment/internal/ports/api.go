package ports

import (
	"context"

	"github.com/GolangAndGrpc/microservices/payment/internal/application/core/domain"
)

type ApiPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}