package ports

import "github.com/mic-x01/grpc-microservices-in-go/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
