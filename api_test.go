package main

import (
	"github.com/selvatico/go-mocket"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCreateCardioWorkout(t *testing.T) {
	router := SetupRouter()
	workout_data := url.Values{}
	workout_data.Add("name", "Run")
	workout_data.Add("type", "cardio")
	workout_data.Add("exp", "5")
	workout_data.Add("distance", "6")
	workout_data.Add("time", "6000")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/workouts/create",
		strings.NewReader(workout_data.Encode()),
	)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Expected 200 but got %s", string(req.Response.StatusCode))
	}

	t.Log("Cardio Workout Created!")
}

func TestGetWorkout(t *testing.T) {
	t.Log("TestGetWorkout")

}

func TestGetWorkouts(t *testing.T) {
	t.Log("TestGetWorkouts")

}

func TestUpdateWorkout(t *testing.T) {
	t.Log("TestUpdateWorkout")

}

func TestDeleteWorkout(t *testing.T) {
	t.Log("TestDeleteWorkout")

}
