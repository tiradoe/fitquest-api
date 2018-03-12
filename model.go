package main

import (
    "github.com/jinzhu/gorm"
)

type Weekday string

const (
    Monday Weekday = "Monday"
    Tuesday Weekday = "Tuesday"
    Wednesday Weekday = "Wednesday"
    Thursday Weekday = "Thursday"
    Saturday Weekday = "Saturday"
    Sunday Weekday = "Sunday"
)


type Routine struct {
    gorm.Model
    Name string `json:"routine_name"`
    Day Weekday `json:"weekday"`
    Workouts []Workout `json:"workouts"`
}


type Workout struct {
    gorm.Model
    RoutineId int
    Name string `json:"workout_name"`
    Type string `json:"workout_type"`
    Sets []Set `json:"sets"`
    Distance int `json:"distance"`
    Time int `json:"time"`
    Weight int `json:"weight"`
    IncreaseBy int `json:"increase_by"`
    Experience int `json:"exp"`
}


type Set struct {
    gorm.Model
    WorkoutId int
    Reps int
    IncreaseWeight bool
}
