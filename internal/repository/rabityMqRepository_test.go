package repository

import (
	"FastGourmet/internal/core/domain"
	"github.com/rabbitmq/amqp091-go"
	"testing"
)

func TestRabittMqRepository_Enqueue(t *testing.T) {
	type fields struct {
		conn *amqp091.Connection
	}
	type args struct {
		order domain.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RabittMqRepository{
				conn: tt.fields.conn,
			}
			if err := r.Enqueue(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRabittMqRepository_Receive(t *testing.T) {
	type fields struct {
		conn *amqp091.Connection
	}
	type args struct {
		order domain.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RabittMqRepository{
				conn: tt.fields.conn,
			}
			if err := r.Receive(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
