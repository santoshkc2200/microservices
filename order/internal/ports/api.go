package ports

import (
	"context"

	"gitHub.com/santoshkc2200/microservices/order/internal/application/core/domain"
)

type APIPorts interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}
