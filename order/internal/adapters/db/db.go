package db

import (
	"fmt"

	"github.com/GolangAndGrpc/microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Adapter struct{
	db *gorm.DB
}

type Order struct{
	gorm.Model
	CustomerID int64
	Status string
	orderItems []OrderItem
}

type OrderItem struct{
	gorm.Model
	ProductCode string
	UnitPrice float32
	Quantity int32
	OrderID uint
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil{
		return nil, fmt.Errorf("db connection error : %v", openErr)
	}

	err := db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db:db}, nil
}


func (a Adapter) Get (id string) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.orderItems{
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: float64(orderItem.UnitPrice),
			Quantity: orderItem.Quantity,
		})
	}

	order := domain.Order{
		ID: int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status: orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt: orderEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: float32(orderItem.UnitPrice),
			Quantity: orderItem.Quantity,
		})
	}
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status: order.Status,
		orderItems: orderItems,
	}

	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}