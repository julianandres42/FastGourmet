package orderService

import (
	"FastGourmet/internal/core/domain"
	"FastGourmet/internal/core/port"
)

type service struct {
	orderStorageRepository port.OrderStorageRepository
	orderQueueRepository   port.OrdersQueueRepository
}

func New() *service {
	return &service{}
}

func (s service) Send(order domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(order domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Get(id int) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(id int, order domain.Order) error {
	//TODO implement me
	panic("implement me")
}
