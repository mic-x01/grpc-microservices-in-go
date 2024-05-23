package ports

import "github.com/mic-x01/grpc-microservices-in-go/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
