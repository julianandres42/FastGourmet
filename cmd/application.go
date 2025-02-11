package cmd

import (
	"FastGourmet/internal/core/domain"
	"FastGourmet/internal/core/port"
	"FastGourmet/internal/core/service/orderService"
	"FastGourmet/internal/repository"
	"context"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	orderService port.OrdersService
}

func (app *Application) Start() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	app.orderService = orderService.NewServiceImp(repository.NewMongoDbRepository(client), repository.NewRabittMqRepository(conn))
}

func (app *Application) EnqueueOrder(c *gin.Context) {
	var newOrder domain.Order
	if err := c.ShouldBind(&newOrder); err == nil {
		app.orderService.Send(&newOrder)
	}
}

func (app *Application) Receive() {
	app.orderService.Receive()
}

func (app *Application) Recover() {
	app.orderService.Recover()
}

func (app *Application) Get(c *gin.Context) {
	res, err := app.orderService.Get(c.Query("id"))
	if err != nil {
		c.JSON(500, gin.H{})
		return
	}
	c.JSON(200, res)
	return
}

func (app *Application) Update(c *gin.Context) {
	var newOrder domain.Order
	if err := c.ShouldBind(&newOrder); err == nil {
		app.orderService.Update(newOrder.ID, &newOrder)
	}
}

func (app *Application) List(c *gin.Context) {
	res, err := app.orderService.List()
	if err != nil {
		c.JSON(500, gin.H{})
	}
	c.JSON(200, res)
}
