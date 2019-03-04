package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

func CreateWorkout(c *gin.Context) {
	var sets []Set

	name := c.PostForm("name")
	workout_type := c.PostForm("type")
	sets = GenerateSets()
	exp, _ := strconv.Atoi(c.PostForm("exp"))
	distance, _ := strconv.Atoi(c.PostForm("distance"))
	time, _ := strconv.Atoi(c.PostForm("time"))
	weight, _ := strconv.Atoi(c.PostForm("weight"))
	increase_by, _ := strconv.Atoi(c.PostForm("increase_by"))

	db := Database()

	if workout_type == "cardio" {
		var cardio Workout

		cardio = Workout{
			Name:       name,
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

	} else if workout_type == "strength" {
		var strength Workout

		strength = Workout{
			Name:       name,
			Type:       workout_type,
			Sets:       sets,
			Weight:     weight,
			IncreaseBy: increase_by,
			Experience: exp,
		}

		db.Create(&strength)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Workout Created!",
			"workout": strength,
		})
	} else {
		db.Close()
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No Workout Found",
		})
	}

	db.Close()

	return
}

func GetWorkout(c *gin.Context) {
	var workout Workout

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("ID not provided")
	}

	db := Database()

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

func GetWorkouts(c *gin.Context) {
	var workouts []Workout

	db := Database()
	db.Preload("Sets").Find(&workouts)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"workouts": workouts,
	})

	return
}

func UpdateWorkout(c *gin.Context) {
	var cardio Workout
	var strength Workout

	id := c.Param("id")
	workout_type := c.Param("type")
	name := c.PostForm("name")
	sets := GenerateSets()
	exp, _ := strconv.Atoi(c.PostForm("exp"))

	db := Database()

	if workout_type == "cardio" {
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
	} else if workout_type == "strength" {
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

func DeleteWorkout(c *gin.Context) {
	var workout Workout

	id := c.Param("id")

	db := Database()

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

func GenerateSets() []Set {
	/* Just for testing.*/
	//@todo delete this when UI is ready
	var set_test []Set
	set_test = append(set_test, Set{Reps: 5, IncreaseWeight: true})
	set_test = append(set_test, Set{Reps: 3, IncreaseWeight: false})

	return set_test
}
