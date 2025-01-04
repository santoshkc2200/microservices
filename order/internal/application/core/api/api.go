package api

import (
	"context"

	"gitHub.com/santoshkc2200/microservices/order/internal/application/core/domain"
	"gitHub.com/santoshkc2200/microservices/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
