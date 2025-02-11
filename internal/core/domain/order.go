package domain

import (
	"time"
)

type Order struct {
	ID          string        `json:"id"`
	ArrivedTime time.Time     `json:"arrivedTime"`
	Dishes      []string      `json:"dishes"`
	Priority    OrderPriority `json:"priority"`
	Status      OrderState    `json:"status"`
	Source      OrderSource   `json:"source"`
}

type OrderState int
type OrderSource int

type OrderPriority int

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

const (
	normal OrderPriority = iota
	hight  OrderPriority = iota
)

var OrderPriorityStrings = map[OrderPriority]string{
	normal: "NORMAL",
	hight:  "HIGH",
}
