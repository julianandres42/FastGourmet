package orderService

import (
	"FastGourmet/internal/core/domain"
	"encoding/json"
	"errors"
	"testing"
	"time"
)

type orderStorageRepositoryMock struct {
}

func (o orderStorageRepositoryMock) Save(order *domain.Order) error {
	//TODO implement me
	return nil
}

func (o orderStorageRepositoryMock) List() ([]*domain.Order, error) {
	//TODO implement me
	return []*domain.Order{{
		ID:          "",
		ArrivedTime: time.Time{},
		Dishes:      nil,
		Priority:    0,
		Status:      0,
		Source:      0,
	}}, nil
}

func (o orderStorageRepositoryMock) Get(id string) (*domain.Order, error) {
	//TODO implement me
	return &domain.Order{}, nil
}

func (o orderStorageRepositoryMock) Update(id string, order *domain.Order) error {
	//TODO implement me
	return nil
}

type orderStorageRepositoryMockError struct{}

func (o orderStorageRepositoryMockError) Save(order *domain.Order) error {
	//TODO implement me
	return errors.New("error")
}

func (o orderStorageRepositoryMockError) List() ([]*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o orderStorageRepositoryMockError) Get(id string) (*domain.Order, error) {
	//TODO implement me
	return &domain.Order{}, nil
}

func (o orderStorageRepositoryMockError) Update(id string, order *domain.Order) error {
	//TODO implement me
	panic("implement me")
}

type orderQueueRepositoryMock struct{}

func (o orderQueueRepositoryMock) Enqueue(order *domain.Order) error {
	return nil
}

func (o orderQueueRepositoryMock) Receive(c chan []byte) {
	c <- []byte("hello world")

}

func TestSendOrder(t *testing.T) {
	service := NewServiceImp(orderStorageRepositoryMock{}, orderQueueRepositoryMock{})
	result := service.Send(&domain.Order{
		ID:          "",
		ArrivedTime: time.Time{},
		Dishes:      nil,
		Priority:    0,
		Status:      0,
		Source:      0,
	})
	if result != nil {
		t.Error("result should be nil")
	}
}

func TestReceiveOrder(t *testing.T) {
	service := NewServiceImp(orderStorageRepositoryMock{}, orderQueueRepositoryMock{})
	go service.Receive()
	result := <-service.messages
	if result == nil {
		t.Error("result should not be nil")
	}
}

func TestRecoveryMessage(t *testing.T) {
	service := NewServiceImp(orderStorageRepositoryMock{}, orderQueueRepositoryMock{})
	order := domain.Order{
		ID:          "",
		ArrivedTime: time.Time{},
		Dishes:      nil,
		Priority:    0,
		Status:      0,
		Source:      0,
	}
	message, _ := json.Marshal(order)
	go func() {
		service.messages <- message
	}()
	go service.Recover()
	if len(service.messages) != 0 {
		t.Error("message should be empty")
	}
}

func TestSaveError(t *testing.T) {
	service := NewServiceImp(orderStorageRepositoryMockError{}, orderQueueRepositoryMock{})
	err := service.Create(&domain.Order{})
	if err == nil {
		t.Error("error should be nil")
	}
}

func TestGetOrder(t *testing.T) {
	service := NewServiceImp(orderStorageRepositoryMock{}, orderQueueRepositoryMock{})
	result, err := service.Get("aaa")
	if err != nil {
		t.Error("error should be nil")
	}
	if result == nil {
		t.Error("result should not be nil")
	}
}

func TestUpdateOrder(t *testing.T) {
	service := NewServiceImp(orderStorageRepositoryMock{}, orderQueueRepositoryMock{})
	err := service.Update("aaa", &domain.Order{})
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestListOrder(t *testing.T) {
	service := NewServiceImp(orderStorageRepositoryMock{}, orderQueueRepositoryMock{})
	result, err := service.List()
	if err != nil {
		t.Error("error should be nil")
	}
	if len(result) == 0 {
		t.Error("result should not be empty")
	}
}
