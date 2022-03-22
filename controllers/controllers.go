package controllers

import (
	"lb-user-management-system/models"
	"net/http"

	"github.com/sirupsen/logrus"

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
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Errorf("%s", err)
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
		logrus.WithFields(logrus.Fields{
			"error": err,
			"id":    ctx.Param("id"),
		}).Errorf("Failed to find the book with id %s", ctx.Param("id"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

//Update the User - PATCH
func UpdateUser(ctx *gin.Context) {
	//get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"id":    ctx.Param("id"),
		}).Errorf("no record found")
		return
	}
	//validate the input
	var input models.UpdateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Errorf("input format incorrect")
		return
	}
	models.DB.Model(&user).Update(input) // update the model with Model()
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

//delete the user
func DeleteUser(ctx *gin.Context) {
	//find the user
	var user models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"id":    ctx.Param("id"),
		}).Errorf("no record found")
		return
	}
	//delete the user from database using delete() provided by middleware
	models.DB.Delete(&user)
	//return deeted user
	ctx.JSON(http.StatusOK, gin.H{"data": true}) //nopoint in returning deleted object
}
