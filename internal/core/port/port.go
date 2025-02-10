package port

import "FastGourmet/internal/core/domain"

type OrdersQueueRepository interface {
	Enqueue(order domain.Order) error
	Receive(order domain.Order) error
}

type OrderStorageRepository interface {
	Save(order domain.Order) error
	List() ([]domain.Order, error)
	Get(id int) (domain.Order, error)
	Update(id int, order domain.Order) error
}

type OrdersService interface {
	Send(order domain.Order) error
	Create(order domain.Order) error
	Get(id int) (domain.Order, error)
	Update(id int, order domain.Order) error
}
