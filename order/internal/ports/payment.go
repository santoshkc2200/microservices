package ports

import "gitHub.com/santoshkc2200/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
