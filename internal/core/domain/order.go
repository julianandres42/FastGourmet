package domain

import "time"

type Order struct {
	ID          string      `json:"id"`
	ArrivedTime time.Time   `json:"arrivedTime"`
	Dishes      []string    `json:"dishes"`
	Status      OrderState  `json:"status"`
	Source      OrderSource `json:"source"`
}

type OrderState int
type OrderSource int

const (
	Pending OrderState = iota
	Preparing
	Ready
	delivered
)

var OrderStateStrings = map[OrderState]string{
	Pending:   "PENDING",
	Preparing: "PREPARING",
	Ready:     "READY",
	delivered: "DELIVERED",
}

const (
	OnSite OrderSource = iota
	Delivery
	Phone
)

var OrderSourceStrings = map[OrderSource]string{
	OnSite:   "ON_SITE",
	Delivery: "DELIVERY",
	Phone:    "PHONE",
}
