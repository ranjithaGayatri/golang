package router

import (
	"PackageApi/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func RoutesDefined() {

	Router.GET("/user/:id", controllers.GetUsersByAge)
	Router.POST("/user", controllers.CreateUser)
	Router.GET("/user", controllers.GetUsers)
	Router.PUT("/users", controllers.UpdateUsers)
	Router.Run("localhost:8085")
}

func IamRoute() {
	fmt.Println("I am route1")
}
