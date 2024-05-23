package ports

import "github.com/mic-x01/grpc-microservices-in-go/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
