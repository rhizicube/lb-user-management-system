package main

import (
	"fmt"
	"lb-user-management-system/controllers"
	"lb-user-management-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var c = make(chan models.CreateUser)

func main() {

	sr := gin.Default() // initialize new Default router within sr variable

	models.ConnectDataBase()

	go createUser(c)

	sr.GET("/users", controllers.FindUsers)
	// sr.POST("/user", controllers.CreateUser)
	sr.POST("/user", createUser1)
	sr.GET("/user/:id", controllers.FindUser)
	sr.PATCH("/user/:id", controllers.UpdateUser) // for update we use PATCH
	sr.DELETE("/user/:id", controllers.DeleteUser)
	sr.Run("0.0.0.0:8000") // run our server
}

func createUser1(ctx *gin.Context) {
	var input models.CreateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c <- input
}

func createUser(c chan models.CreateUser) {
	for {
		input := <-c
		user := models.User{Name: input.Name, Contact: input.Contact, Address: input.Address}
		models.DB.Create(&user)
		fmt.Printf("USER IS SUCCESSFULLY CREATED")
	}
}
