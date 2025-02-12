package cmd

import (
	"FastGourmet/internal/core/domain"
	"testing"
)

type serviceMock struct {
}

func (s serviceMock) Send(order *domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceMock) Create(order *domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceMock) Get(id string) (*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceMock) Update(id string, order *domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceMock) Receive() {
	//TODO implement me
	panic("implement me")
}

func (s serviceMock) Recover() {
	//TODO implement me
	panic("implement me")
}

func (s serviceMock) List() ([]*domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func TestApplication_Start(t *testing.T) {

}
