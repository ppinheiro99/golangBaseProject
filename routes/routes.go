package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golangBaseProject/controller"
)

func Setup(r *gin.Engine){
	r.GET("/ping", controller.Example)
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
}