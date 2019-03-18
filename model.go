package main

import (
	"github.com/jinzhu/gorm"
)

// Weekday stores a day of the week
type Weekday string

// Days of the week
const (
	Monday    Weekday = "Monday"
	Tuesday   Weekday = "Tuesday"
	Wednesday Weekday = "Wednesday"
	Thursday  Weekday = "Thursday"
	Saturday  Weekday = "Saturday"
	Sunday    Weekday = "Sunday"
)

// Routine stores workouts that represent a daily routine
type Routine struct {
	gorm.Model
	Name     string    `json:"routine_name"`
	Day      Weekday   `json:"weekday"`
	Workouts []Workout `json:"workouts"`
}

// Workout stores details for a performed workout
type Workout struct {
	gorm.Model
	RoutineID  int
	Name       string `json:"workout_name"`
	Type       string `json:"workout_type"`
	Sets       []Set  `json:"sets"`
	Distance   int    `json:"distance"`
	Time       int    `json:"time"`
	Weight     int    `json:"weight"`
	IncreaseBy int    `json:"increase_by"`
	Experience int    `json:"exp"`
}

// Set stores groups of workouts
type Set struct {
	gorm.Model
	WorkoutID      int
	Reps           int
	IncreaseWeight bool
}
