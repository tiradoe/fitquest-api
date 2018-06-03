package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	db := Database()
	db.AutoMigrate(&Routine{}, &Workout{}, &Set{})

	router := gin.Default()

	v1 := router.Group("/api/v1/workouts")
	{
		v1.GET("/", GetWorkouts)
		v1.GET("/:id", GetWorkout)
		v1.PUT("/:id/:type", UpdateWorkout)
		v1.POST("/create", CreateWorkout)
		v1.DELETE("/:id/:type", DeleteWorkout)
	}
	router.Run()
}
