package orderService

import (
	"FastGourmet/internal/core/domain"
	"FastGourmet/internal/core/port"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Create(t *testing.T) {
	type fields struct {
		orderStorageRepository port.OrderStorageRepository
		orderQueueRepository   port.OrdersQueueRepository
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
			s := service{
				orderStorageRepository: tt.fields.orderStorageRepository,
				orderQueueRepository:   tt.fields.orderQueueRepository,
			}
			if err := s.Create(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Get(t *testing.T) {
	type fields struct {
		orderStorageRepository port.OrderStorageRepository
		orderQueueRepository   port.OrdersQueueRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				orderStorageRepository: tt.fields.orderStorageRepository,
				orderQueueRepository:   tt.fields.orderQueueRepository,
			}
			got, err := s.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Send(t *testing.T) {
	type fields struct {
		orderStorageRepository port.OrderStorageRepository
		orderQueueRepository   port.OrdersQueueRepository
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
			s := service{
				orderStorageRepository: tt.fields.orderStorageRepository,
				orderQueueRepository:   tt.fields.orderQueueRepository,
			}
			if err := s.Send(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Update(t *testing.T) {
	type fields struct {
		orderStorageRepository port.OrderStorageRepository
		orderQueueRepository   port.OrdersQueueRepository
	}
	type args struct {
		id    int
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
			s := service{
				orderStorageRepository: tt.fields.orderStorageRepository,
				orderQueueRepository:   tt.fields.orderQueueRepository,
			}
			if err := s.Update(tt.args.id, tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
