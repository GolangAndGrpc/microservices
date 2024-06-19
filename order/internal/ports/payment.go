package ports

import "github.com/GolangAndGrpc/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}