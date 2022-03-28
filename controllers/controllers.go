package controllers

import (
	"fmt"
	"lb-user-management-system/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var wg = sync.WaitGroup{} // added wait group

//Get list of all users
func FindUsers(ctx *gin.Context) {
	wg.Add(1)
	go FindUsersHelper(ctx)
	fmt.Printf("Data will be fetched soon")
	wg.Wait()
}
func FindUsersHelper(ctx *gin.Context) {
	time.Sleep(5 * time.Second) // let request take 5 seconds to give output
	var users []models.User
	models.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})
	wg.Done()
}

// Create User using Create()
func CreateUser(ctx *gin.Context) {
	wg.Add(1)
	go CreateUserHelper(ctx)
	fmt.Printf("Data will be created soon")
	wg.Wait()
}
func CreateUserHelper(ctx *gin.Context) {
	time.Sleep(5 * time.Second) // let request take 5 seconds to give output
	//validate User
	var input models.CreateUser
	if err := ctx.ShouldBindJSON(&input); err != nil { // syntax??
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		wg.Done()
		return
	}
	//create user
	user := models.User{Name: input.Name, Contact: input.Contact, Address: input.Address}
	models.DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"data": user})
	wg.Done()
}

//GET user using Id i.e. single user
func FindUser(ctx *gin.Context) {
	wg.Add(1)
	go FindUserHelper(ctx)
	fmt.Printf("Data will be fetched soon")
	wg.Wait()
}
func FindUserHelper(ctx *gin.Context) {
	time.Sleep(5 * time.Second) // let request take 5 seconds to give output
	var book models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		wg.Done()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
	wg.Done()
}

//Update the User - PATCH
func UpdateUser(ctx *gin.Context) {
	wg.Add(1)
	go UpdateUserHelper(ctx)
	fmt.Printf("Data will be updated soon")
	wg.Wait()
}
func UpdateUserHelper(ctx *gin.Context) {
	time.Sleep(5 * time.Second) // let request take 5 seconds to give output
	//get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no record found"})
		wg.Done()
		return
	}
	//validate theinput
	var input models.UpdateUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		wg.Done()
		return
	}
	models.DB.Model(&user).Update(input) // update the model with Model()
	ctx.JSON(http.StatusOK, gin.H{"data": user})
	wg.Done()
}

//delete the user
func DeleteUser(ctx *gin.Context) {
	wg.Add(1)
	go DeleteUserHelper(ctx)
	fmt.Printf("Data will be deleted soon")
	wg.Wait()
}
func DeleteUserHelper(ctx *gin.Context) {
	time.Sleep(5 * time.Second) // let request take 5 seconds to give output
	//find the user
	var user models.User
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		wg.Done()
		return
	}
	//delete the user from database using delete() provided by middleware
	models.DB.Delete(&user)
	//return deeted user
	ctx.JSON(http.StatusOK, gin.H{"data": true}) //nopoint in returning deleted object
	wg.Done()
}
