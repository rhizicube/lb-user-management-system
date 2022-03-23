package main

import (
	"lb-user-management-system/controllers"
	"lb-user-management-system/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel) //setting level to debug

	sr := gin.Default() // initialize new Default router within sr variable

	models.ConnectDataBase()

	sr.GET("/users", controllers.FindUsers)
	sr.POST("/user", controllers.CreateUser)
	sr.GET("/user/:id", controllers.FindUser)
	sr.PATCH("/user/:id", controllers.UpdateUser) // for update we use PATCH
	sr.DELETE("/user/:id", controllers.DeleteUser)
	sr.Run("0.0.0.0:8000") // run our server
}
