package ports

import (
	"context"

	"gitHub.com/santoshkc2200/microservices/order/internal/application/core/domain"
)

type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(context.Context, *domain.Order) error
}
