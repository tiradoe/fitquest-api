package main

func main() {
	db := Database()
	db.AutoMigrate(&Routine{}, &Workout{}, &Set{})

	router := SetupRouter()
	router.Run()
}
