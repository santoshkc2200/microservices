package api

import (
	"context"

	"gitHub.com/santoshkc2200/microservices/order/internal/application/core/domain"
	"gitHub.com/santoshkc2200/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		// Resolves status from a payment error
		st, _ := status.FromError(paymentErr)
		//Payment error as a separate field
		fieldError := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: st.Message(),
		}
		// Initiates a bad request error
		badReq := &errdetails.BadRequest{}
		// populates with the actual payment details
		badReq.FieldViolations = append(badReq.FieldViolations, fieldError)
		// creates the root status
		orderStatus := status.New(codes.InvalidArgument, "order creation fialed")
		// Augments the status with a payment error
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetails.Err()
	}
	return order, nil
}
