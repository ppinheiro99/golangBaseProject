package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golangBaseProject/backEnd/model"
	"github.com/golangBaseProject/backEnd/services"
	"golang.org/x/crypto/bcrypt"
)

func Example(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

const SecretKey = "secret"

func Register(c *gin.Context){
	var usr model.Users

	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"}) // Ver melhor estes "prints" de erros
		return
	}
	if len(usr.Email) < 8 || len(usr.Email) > 254 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Email"})
		return	
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(usr.Password), 14)
	usr.Password = string(password)
	services.Db.Create(&usr)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "User ID": usr.ID, "user": usr})
}

func Login(c *gin.Context){
	var usr model.Users
	var creds model.Users

	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	services.Db.Find(&creds, "email = ?", usr.Email)

	if creds.ID == 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"}) 
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(creds.Password),[]byte(usr.Password)); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Incorrect Password!"}) 
		return
	}

	// Create token 
	token := services.GenerateTokenJWT(creds)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Access denied!"})
		return
	}

	// Return Token
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "token": token, "email": creds.Email})

}