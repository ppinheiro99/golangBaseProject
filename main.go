package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golangBaseProject/services"
	"github.com/golangBaseProject/routes"
)


func init() {
	services.OpenDatabase()
}

func main() {
	r := gin.Default()
	routes.Setup(r)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}