package orderService

import (
	"FastGourmet/internal/core/domain"
	"FastGourmet/internal/core/port"
	"encoding/json"
	"github.com/google/uuid"
	"log"
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
	id := uuid.New()
	order.ID = string(id.String())
	order.ArrivedTime = time.Now()
	order.Priority = 0
	return s.orderQueueRepository.Enqueue(order)
}

func (s serviceImp) Receive() {
	s.orderQueueRepository.Receive(s.messages)
}

func (s serviceImp) Recover() {
	for {
		message := <-s.messages
		log.Printf("Got it again [x] %s", message)
		var newOrder = domain.Order{}
		json.Unmarshal(message, &newOrder)
		s.Create(&newOrder)
	}
}

func (s serviceImp) Create(order *domain.Order) error {
	err := s.orderStorageRepository.Save(order)
	if err != nil {
		return err
	}
	return nil
}

func (s serviceImp) Get(id string) (*domain.Order, error) {
	//TODO implement me
	return s.orderStorageRepository.Get(id)
}

func (s serviceImp) Update(id string, order *domain.Order) error {
	//TODO implement me
	return s.orderStorageRepository.Update(id, order)
}

func (s serviceImp) List() ([]*domain.Order, error) {
	return s.orderStorageRepository.List()
}
