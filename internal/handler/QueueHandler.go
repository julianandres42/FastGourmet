package handler

import (
	"FastGourmet/internal/core/domain"
	"FastGourmet/internal/core/port"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPHandler struct {
	orderService port.OrdersService
}

func NewHTTPHandler(orderService port.OrdersService) *HTTPHandler {
	return &HTTPHandler{
		orderService: orderService,
	}
}

func (hdl *HTTPHandler) Post(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := hdl.orderService.Send(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}
