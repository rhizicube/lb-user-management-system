package controllers

import (
	"lb-user-management-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get list of all users
func FindUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Create User using Create()
func CreateUser(ctx *gin.Context) {
	//validate User
	var input models.CreateUser
	if err := ctx.ShouldBindJSON(&input); err != nil { // syntax??
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//create user
	user := models.User{Name: input.Name, Contact: input.Contact, Address: input.Address}
	models.DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

//GET user using Id i.e. single user
func FindUser(ctx *gin.Context) {
	var book models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

//Update the User - PATCH
func UpdateUser(ctx *gin.Context) {
	//get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no record found"})
		return
	}
	//validate theinput
	var input models.UpdateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// models.DB.Model(&user).Update(input)
	// update the model with Model()
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

//delete the user
func DeleteUser(ctx *gin.Context) {
	//find the user
	var user models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//delete the user from database using delete() provided by middleware
	models.DB.Delete(&user)
	//return deeted user
	ctx.JSON(http.StatusOK, gin.H{"data": true}) //nopoint in returning deleted object
}
