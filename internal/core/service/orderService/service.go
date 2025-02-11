package orderService

import (
	"FastGourmet/internal/core/domain"
	"FastGourmet/internal/core/port"
	"os/exec"
	"time"
)

type serviceImp struct {
	orderStorageRepository port.OrderStorageRepository
	orderQueueRepository   port.OrdersQueueRepository
	messages               chan []byte
}

func NewServiceImp(orderStorageRepository port.OrderStorageRepository, orderQueueRepository port.OrdersQueueRepository) *serviceImp {
	return &serviceImp{
		orderStorageRepository: orderStorageRepository,
		orderQueueRepository:   orderQueueRepository,
		messages:               make(chan []byte),
	}
}

func (s serviceImp) Send(order *domain.Order) error {
	newUUID, _ := exec.Command("uuidgen").Output()
	order.ID = string(newUUID)
	order.ArrivedTime = time.Now()
	order.Priority = 0
	return s.orderQueueRepository.Enqueue(order)
}

func (s serviceImp) Receive() {
	s.orderQueueRepository.Receive(s.messages)
}

func (s serviceImp) Create(order *domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceImp) Get(id int) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceImp) Update(id int, order *domain.Order) error {
	//TODO implement me
	panic("implement me")
}
