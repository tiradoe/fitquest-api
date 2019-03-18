package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

// CreateWorkout adds a workout to the database
func CreateWorkout(c *gin.Context) {
	var sets []Set

	name := c.PostForm("name")
	workoutType := c.PostForm("type")
	sets = GenerateSets()
	exp, _ := strconv.Atoi(c.PostForm("exp"))
	distance, _ := strconv.Atoi(c.PostForm("distance"))
	time, _ := strconv.Atoi(c.PostForm("time"))
	weight, _ := strconv.Atoi(c.PostForm("weight"))
	increaseBy, _ := strconv.Atoi(c.PostForm("increaseBy"))

	db := getDB()
	defer db.Close()

	if workoutType == "cardio" {
		var cardio Workout

		cardio = Workout{
			Name:       name,
			Type:       workoutType,
			Sets:       sets,
			Distance:   distance,
			Time:       time,
			Experience: exp,
		}
		db.Create(&cardio)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Workout Created!",
			"workout": cardio,
		})

	} else if workoutType == "strength" {
		var strength Workout

		strength = Workout{
			Name:       name,
			Type:       workoutType,
			Sets:       sets,
			Weight:     weight,
			IncreaseBy: increaseBy,
			Experience: exp,
		}

		db.Create(&strength)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Workout Created!",
			"workout": strength,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Workout Found",
		})
	}

	return
}

// GetWorkout gets a single workout using the id
func GetWorkout(c *gin.Context) {
	var workout Workout

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("ID not provided")
	}

	db := getDB()
	defer db.Close()

	if db.Preload("Sets").First(&workout, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"workout": "No Workout Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"workout": workout,
	})

	return
}

// GetWorkouts gets all of the user's workouts
func GetWorkouts(c *gin.Context) {
	var workouts []Workout

	db := getDB()
	defer db.Close()

	db.Preload("Sets").Find(&workouts)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"workouts": workouts,
	})

	return
}

// UpdateWorkout makes user provided changes to the specified workout
func UpdateWorkout(c *gin.Context) {
	var cardio Workout
	var strength Workout

	id := c.Param("id")
	workoutType := c.Param("type")
	name := c.PostForm("name")
	sets := GenerateSets()
	exp, _ := strconv.Atoi(c.PostForm("exp"))

	db := getDB()
	defer db.Close()

	if workoutType == "cardio" {
		distance, _ := strconv.Atoi(c.PostForm("distance"))
		time, _ := strconv.Atoi(c.PostForm("time"))

		if db.First(&cardio, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"workout": "No Workout Found",
			})

			return
		}

		cardio.Name = name
		cardio.Sets = sets
		cardio.Experience = exp
		cardio.Distance = distance
		cardio.Time = time

		db.Save(&cardio)

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"workout": cardio,
			"message": "Workout saved!",
		})
	} else if workoutType == "strength" {
		weight, _ := strconv.Atoi(c.PostForm("weight"))

		if db.First(&strength, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"workout": "No Workout Found",
			})

			return
		}

		strength.Name = name
		strength.Sets = sets
		strength.Weight = weight

		db.Save(&strength)

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"workout": strength,
			"message": "Workout Saved!",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Workout Found",
		})
	}

	return
}

// DeleteWorkout removes the specified workout from the database
func DeleteWorkout(c *gin.Context) {
	var workout Workout

	id := c.Param("id")

	db := getDB()
	defer db.Close()

	if db.First(&workout, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"workout": "No Workout Found",
		})

		return
	}
	db.Delete(&workout)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"workout": "Workout Deleted!",
	})

	return
}

// GenerateSets creates dummy sets for testing.  Should be deleted when UI is ready
func GenerateSets() []Set {
	var setTest []Set
	setTest = append(setTest, Set{Reps: 5, IncreaseWeight: true})
	setTest = append(setTest, Set{Reps: 3, IncreaseWeight: false})

	return setTest
}
