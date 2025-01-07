package payment

import (
	"context"

	"gitHub.com/santoshkc2200/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices-proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerId,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	// Data model for connection configurations
	var opts []grpc.DialOption
	// disabling TLS
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// connects to the service
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}
