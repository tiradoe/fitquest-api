package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)


func main() {
    db := Database()
    db.AutoMigrate(&Workout{})

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

//func CreateTodo(c *gin.Context) {
//    completed, _ := strconv.Atoi(c.PostForm("completed"))
//    todo := Todo{Title: c.PostForm("title"), Completed: completed};
//https://github.com/jinzhu/gorm    db := Database()
//    db.Save(&todo)
//    c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Todo item created successfully!", "resourceId": todo.ID})
//}


func CreateWorkout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Create Workout"})
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
