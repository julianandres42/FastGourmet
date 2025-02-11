package port

import "FastGourmet/internal/core/domain"

type OrdersQueueRepository interface {
	Enqueue(order *domain.Order) error
	Receive(chan []byte)
}

type OrderStorageRepository interface {
	Save(order *domain.Order) error
	List() ([]*domain.Order, error)
	Get(id string) (*domain.Order, error)
	Update(id string, order *domain.Order) error
}

type OrdersService interface {
	Send(order *domain.Order) error
	Create(order *domain.Order) error
	Get(id string) (*domain.Order, error)
	Update(id string, order *domain.Order) error
	Receive()
	Recover()
	List() ([]*domain.Order, error)
}
