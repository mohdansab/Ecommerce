package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohdansab/controller"
	"github.com/mohdansab/db"
	"github.com/mohdansab/middleware"
)

func main() {
	db.ConnectDB()

	r := gin.Default()
	r.GET("/H", middleware.UserAuth(), controller.HelloWorld)
	r.POST("/create", middleware.UserAuth(), controller.Create)
	r.PATCH("/update", middleware.UserAuth(), controller.Update)
	r.DELETE("/delete", middleware.UserAuth(), controller.Delete)
	r.GET("/view", middleware.UserAuth(), controller.ViewAll)
	r.POST("/signup", controller.SignUp)
	r.GET("/login", controller.Login)
	r.POST("/cart", controller.AddCart)
	r.Run(":8083")
}
