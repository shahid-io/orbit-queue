package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shahid-io/orbit-queue/scheduler"
	"net/http"
	"fmt"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/jobs", func(c *gin.Context){
		fmt.Println("c : ",c)
		var job scheduler.JobRequest
		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := scheduler.RegisterJob(job); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register job"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Job scheduled"})
	})
}