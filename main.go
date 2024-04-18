package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goropencho/golang-gin-auth/controllers"
	"github.com/goropencho/golang-gin-auth/initializer"
	"github.com/goropencho/golang-gin-auth/middlewares"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
	initializer.AutoMigrate()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.SignUp)
	r.POST("/", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run(":3000")
}
