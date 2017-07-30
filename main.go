package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)


func main() {
    db := Database()
    db.AutoMigrate(&Workout{}, &Strength{}, &Cardio{})

    router := gin.Default()

    v1 := router.Group("/api/v1/workouts")
    {
        v1.POST("/", CreateWorkout)
        v1.GET("/", GetWorkouts)
        v1.GET("/:id", GetWorkout)
        v1.PUT("/:id", UpdateWorkout)
        v1.DELETE("/:id", DeleteWorkout)
    }
    router.Run()
}


func CreateWorkout(c *gin.Context) {
    db := Database()

    name := c.PostForm("name")
    set,_ := strconv.Atoi(c.PostForm("set"))
    exp,_ := strconv.Atoi(c.PostForm("exp"))

    workout := Workout{Name: name, Set: set, Experience: exp}
    db.Create(&workout)

    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Workout Created!"})
    return
}


func GetWorkout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Get Workout"})
    return
}


func GetWorkouts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Get Workouts"})
    return
}


func UpdateWorkout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Update Workout"})
    return
}


func DeleteWorkout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Delete Workout"})
    return
}
