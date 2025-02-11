package main

import (
	"FastGourmet/cmd"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	app := cmd.Application{}
	app.Start()
	go app.Receive()
	router := gin.Default()
	router.POST("/enqueue", app.EnqueueOrder)
	router.Run(":8080")

}
