package main

import (
	"encoding/json"
	mocket "github.com/selvatico/go-mocket"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type APIResponse struct {
	Message string  `json:"message"`
	Status  int     `json:"status"`
	Workout Workout `json:"workout"`
}

func TestApiRoutes(t *testing.T) {
	var response APIResponse
	router := SetupRouter()

	t.Run("Creating new workout", func(t *testing.T) {
		workoutData := url.Values{}
		workoutData.Set("name", "Test")
		workoutData.Set("type", "cardio")
		workoutData.Set("exp", "5")
		workoutData.Set("distance", "6")
		workoutData.Set("time", "6000")

		w := httptest.NewRecorder()

		req, err := http.NewRequest(
			"POST",
			"/api/v1/workouts/create",
			strings.NewReader(workoutData.Encode()),
		)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(w, req)
		json.Unmarshal(w.Body.Bytes(), &response)

		//Code in response and json should both be 200
		if w.Code != 200 || response.Status != 200 {
			t.Fatalf("Expected 200 but got %s", string(w.Code))
		}

		//Check the fields
		if response.Workout.Name != "Test" {
			t.Fatalf("Expected Test but got %s", response.Workout.Name)
		}
		if response.Workout.Type != "cardio" {
			t.Fatalf("Expected cardio but got %s", response.Workout.Type)
		}
		if response.Workout.Experience != 5 {
			t.Fatalf("Expected 5 but got %d", response.Workout.Experience)
		}
		if response.Workout.Distance != 6 {
			t.Fatalf("Expected 6 but got %d", response.Workout.Distance)
		}
		if response.Workout.Time != 6000 {
			t.Fatalf("Expected 6000 but got %d", response.Workout.Time)
		}
	})

	t.Run("Getting all workouts from database", func(t *testing.T) {
		mocket.Catcher.Logging = true
		w := httptest.NewRecorder()

		commonReply := []map[string]interface{}{{
			"name":       "Test_workout",
			"type":       "cardio",
			"sets":       GenerateSets(),
			"distance":   5,
			"time":       7000,
			"experience": 10,
		}}

		mocket.Catcher.Reset().NewMock().WithQuery(`INSERT INTO "workouts" ("created_at","updated_at","deleted_at","routine_id","name","type","distance","time","weight","increase_by","experience"`).WithReply(commonReply)

		req, err := http.NewRequest(
			"GET",
			"/api/v1/workouts/",
			nil,
		)

		if err != nil {
			t.Fatal(err)
		}

		router.ServeHTTP(w, req)
		log.Println("Workout body", w.Body)

		//Code in response and json should both be 200
		if w.Code != 200 {
			t.Fatalf("Expected 200 but got %s", string(w.Code))
		}
	})
}

func TestGetWorkout(t *testing.T) {
	t.Log("TestGetWorkouts")

}

func TestUpdateWorkout(t *testing.T) {
	t.Log("TestUpdateWorkout")

}

func TestDeleteWorkout(t *testing.T) {
	t.Log("TestDeleteWorkout")

}
