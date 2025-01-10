package controllers

import (
	"PackageApi/models"
	"PackageApi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users, err := services.Service4func()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error())
	}
	context.IndentedJSON(http.StatusOK, users)
}

func UpdateUsers(context *gin.Context) {
	id := context.Param("id")
	var user, NewUser models.User
	err := services.Service3func(&user, id)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error)
	}
	if err := context.ShouldBindJSON(&NewUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if err := models.Db.Model(&user).Updates(NewUser).Error; err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusNotFound, user)
}

func GetUsersByAge(context *gin.Context) {
	id := context.Param("id")
	var user models.User
	err := services.Servicefunc2(&user, id)
	if err != nil {
		context.IndentedJSON(http.StatusOK, err.Error())
	}
	context.IndentedJSON(http.StatusOK, user)
}

func CreateUser(context *gin.Context) {
	var NewUser models.User
	if err := context.ShouldBindJSON(&NewUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	err := services.Servicefunc1(&NewUser)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err.Error)
		return
	}
	context.IndentedJSON(http.StatusOK, NewUser)
}
