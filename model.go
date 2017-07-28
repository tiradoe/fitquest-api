package main

import (
    "github.com/jinzhu/gorm"
)


type Workout struct {
    gorm.Model
    Name string `json:"name"`
    Set int `json:"set"`
    Experience int `json:"exp"`
}


type Cardio struct {
    Workout
    Distance int `json:"distance"`
    Time int `json:"time"`
}


type Strength struct {
    Workout
    Weight int `json:"weight"`
    Reps int `json:"reps"`
}
