package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golangBaseProject/backEnd/controller"
	"github.com/golangBaseProject/backEnd/services"
)

func Setup(r *gin.Engine){
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// Allow all headers, resolve all cors issues
	r.Use(services.GinMiddleware("*"))

	// NO ROUTE
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// AUTH
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}
	
	// HOME 
	home := r.Group("/api")
	home.Use(services.AuthorizationRequired())
	{
		home.GET("/ping", controller.Example)
		home.GET("/getAllUsers", controller.GetAllUsers)
		home.DELETE("/delete/:id", controller.DeleteUser)
	}
}