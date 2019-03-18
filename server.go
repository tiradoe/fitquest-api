package main

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter returns a routergroup
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1/workouts")
	{
		v1.GET("/", GetWorkouts)
		v1.GET("/:id", GetWorkout)
		v1.PUT("/:id/", UpdateWorkout)
		v1.POST("/create", CreateWorkout)
		v1.DELETE("/:id/", DeleteWorkout)
	}

	return router
}
