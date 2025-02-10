package repository

import "FastGourmet/internal/core/domain"

type RabityMqRepository struct {
}

func (r RabityMqRepository) Enqueue(order domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r RabityMqRepository) Receive(order domain.Order) error {
	//TODO implement me
	panic("implement me")
}
