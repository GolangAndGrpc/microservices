package api

import (
	"log"

	"github.com/GolangAndGrpc/microservices/order/internal/application/core/domain"
	"github.com/GolangAndGrpc/microservices/order/internal/ports"
)

type Application struct{
	db ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db: db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil{
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		log.Fatalf("Error while creatig payment....... %v", err)
	}
	return order, nil
}