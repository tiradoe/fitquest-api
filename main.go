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
        v1.PUT("/:id/:type", UpdateWorkout)
        v1.POST("/create", CreateWorkout)
        v1.DELETE("/:id/:type", DeleteWorkout)
    }
    router.Run()
}


func CreateWorkout(c *gin.Context) {
    name := c.PostForm("name")
    workout_type := c.PostForm("type")
    set,_ := strconv.Atoi(c.PostForm("set"))
    exp,_ := strconv.Atoi(c.PostForm("exp"))
    distance,_ := strconv.Atoi(c.PostForm("distance"))
    time,_ := strconv.Atoi(c.PostForm("time"))
    weight,_ := strconv.Atoi(c.PostForm("weight"))
    reps,_ := strconv.Atoi(c.PostForm("reps"))

    db := Database()

    if workout_type == "cardio" {
        var cardio Cardio

        cardio = Cardio{
            Workout: Workout{Name: name, Set: set, Experience: exp},
            Distance: distance,
            Time: time,
        }
        db.Create(&cardio)

    } else if workout_type == "strength" {
        var strength Strength

        strength = Strength {
            Workout: Workout{Name: name, Set: set, Experience: exp},
            Weight: weight,
            Reps: reps,
        }

        db.Create(&strength)
    } else {
        db.Close()
        c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "message": "No Workout Found"})
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
        log.Println("Id not provided")
    }

    db := Database()

    if workout_type == "cardio" {
        if db.First(&cardio,id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "workout":cardio})
    } else if workout_type == "strength" {
        if db.First(&strength,id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "workout":strength})
    } else {
        c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "message": "No Workout Found"})
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
    var cardio Cardio
    var strength Strength

    id := c.Param("id")
    workout_type := c.Param("type")
    name := c.PostForm("name")
    set,_ := strconv.Atoi(c.PostForm("set"))
    exp,_ := strconv.Atoi(c.PostForm("exp"))

    db := Database()

    if workout_type == "cardio" {
        distance,_ := strconv.Atoi(c.PostForm("distance"))
        time,_ := strconv.Atoi(c.PostForm("time"))

        if db.First(&cardio,id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }

        cardio.Name = name
        cardio.Set = set
        cardio.Experience = exp
        cardio.Distance = distance
        cardio.Time = time

        db.Save(&cardio)

        c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "workout":cardio, "message": "Workout saved!"})
    } else if workout_type == "strength" {
        weight,_ := strconv.Atoi(c.PostForm("weight"))
        reps,_ := strconv.Atoi(c.PostForm("reps"))

        if db.First(&strength,id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }

        strength.Name = name
        strength.Set = set
        strength.Weight = weight
        strength.Reps = reps

        db.Save(&strength)

        c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "workout":strength, "message": "Workout Saved!"})
    } else {
        c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "message": "No Workout Found"})
    }

    return
}


func DeleteWorkout(c *gin.Context) {
    var cardio Cardio
    var strength Strength

    id := c.Param("id")
    workout_type := c.Param("type")

    db := Database()

    if workout_type == "cardio" {
        if db.First(&cardio, id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }
        db.Delete(&cardio)
    } else if workout_type == "strength" {
        if db.First(&strength, id).RecordNotFound() {
            c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "workout":"No Workout Found"})
            return
        }
        db.Delete(&strength)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"status":http.StatusNotFound, "message": "No Workout Found"})
    }

    c.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "workout":"Workout Deleted!"})
    return
}
