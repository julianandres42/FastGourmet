package repository

import "FastGourmet/internal/core/domain"

type RabittMqRepository struct {
}

func (r RabittMqRepository) Enqueue(order domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r RabittMqRepository) Receive(order domain.Order) error {
	//TODO implement me
	panic("implement me")
}
