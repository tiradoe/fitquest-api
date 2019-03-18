package main

func main() {
	db := Database()

	db.AutoMigrate(&Routine{}, &Workout{}, &Set{})
	db.Close()

	router := SetupRouter()
	router.Run()
}
