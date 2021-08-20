package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangBaseProject/backEnd/model"
	"github.com/golangBaseProject/backEnd/services"
)


func GetAllUsers(c *gin.Context) {
	var users []model.Users
	services.Db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": users})
}