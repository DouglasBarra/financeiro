package main

import (
	"backend/controllers"
	"backend/data"
	"github.com/gin-gonic/gin"
)

func main() {
	data.ConnectDatabase()

	r := gin.Default()

	r.POST("index", controllers.PostBank)
	r.GET("/index", controllers.GetAll)

	r.Run(":8082")

}
