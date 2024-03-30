package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdansab/controller"
	"github.com/mohdansab/db"
)

func main() {
	db.ConnectDB()

	r := gin.Default()
	r.GET("/H", controller.HelloWorld)
	r.POST("/create", controller.Create)
	r.PATCH("/update", controller.Update)
	r.DELETE("/delete", controller.Delete)
	r.GET("/view", controller.ViewAll)
	r.POST("/signup", controller.SignUp)
	r.GET("/login", controller.Login)
	r.Run(":8083")
}
