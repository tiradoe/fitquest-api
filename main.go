package main

import (
    "net/http"
    "strconv"
    "log"

    "github.com/gin-gonic/gin"
)


func main() {
    db := Database()
    db.AutoMigrate(&Strength{}, &Cardio{})

    router := gin.Default()

    v1 := router.Group("/api/v1/workouts")
    {
        v1.GET("/", GetWorkouts)
        v1.GET("/:id/:type", GetWorkout)
        v1.PUT("/:id", UpdateWorkout)
        v1.POST("/create", CreateWorkout)
        v1.DELETE("/:id", DeleteWorkout)
    }
    router.Run()
}


func CreateWorkout(c *gin.Context) {
    db := Database()

    name := c.PostForm("name")
    workout_type := c.PostForm("type")
    set,_ := strconv.Atoi(c.PostForm("set"))
    exp,_ := strconv.Atoi(c.PostForm("exp"))
    distance,_ := strconv.Atoi(c.PostForm("distance"))
    time,_ := strconv.Atoi(c.PostForm("time"))
    weight,_ := strconv.Atoi(c.PostForm("weight"))
    reps,_ := strconv.Atoi(c.PostForm("reps"))

    if (workout_type == "cardio") {
        var cardio Cardio

        cardio = Cardio{
            Workout: Workout{Name: name, Set: set, Experience: exp},
            Distance: distance,
            Time: time,
        }
        db.Create(&cardio)
    } else {
        var strength Strength

        strength = Strength {
            Workout: Workout{Name: name, Set: set, Experience: exp},
            Weight: weight,
            Reps: reps,
        }
        db.Create(&strength)
    }

    db.Close()

    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "Workout Created!"})
    return
}


func GetWorkout(c *gin.Context) {
    var cardio Cardio
    var strength Strength
    workout_type := c.Param("type")
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Fatal(err)
    }

    db := Database()

    if(workout_type == "cardio"){
        if db.First(&cardio,id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "workout":cardio})
    } else {
        if db.First(&strength,id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "workout":strength})
    }

    return
}


func GetWorkouts(c *gin.Context) {
    var cardios []Cardio
    var strengths []Strength

    db := Database()
    db.Find(&cardios)
    db.Find(&strengths)

    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "cardio": cardios, "strength": strengths})
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
