package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func StartServer() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong",
		})
	})
	
	RegisterRoutes(router)
	router.Run()
}
