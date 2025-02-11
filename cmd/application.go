package cmd

import (
	"FastGourmet/internal/core/domain"
	"FastGourmet/internal/core/port"
	"FastGourmet/internal/core/service/orderService"
	"FastGourmet/internal/repository"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Application struct {
	orderService port.OrdersService
}

func (app *Application) Start() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	app.orderService = orderService.NewServiceImp(nil, repository.NewRabittMqRepository(conn))
}

func (app *Application) EnqueueOrder(c *gin.Context) {
	var newOrder domain.Order
	if err := c.ShouldBind(&newOrder); err == nil {
		app.orderService.Send(&newOrder)
	}
}

func (app *Application) Receive() {}
