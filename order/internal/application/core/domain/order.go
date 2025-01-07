package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Order struct {
	ID         int64       `json:"id"`
	CustomerId int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, OrderItem := range o.OrderItems {
		totalPrice += OrderItem.UnitPrice * float32(OrderItem.Quantity)
	}
	return totalPrice
}