package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	mocket "github.com/selvatico/go-mocket"
)

func getDB() *gorm.DB {
	isTesting := false

	if flag.Lookup("test.v") != nil {
		isTesting = true
	}
	var db *gorm.DB

	if isTesting {
		db = TestDatabase()
	} else {
		db = Database()
	}

	return db
}

func Database() *gorm.DB {
	var login string = os.Getenv("FITQUESTDB")
	var db_conn string = fmt.Sprintf("%s@tcp(fitquestdb:3306)/fitquest?charset=utf8&parseTime=True&loc=Local", login)

	db, err := gorm.Open("mysql", db_conn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestDatabase() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, err := gorm.Open(mocket.DriverName, "")
	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}

	return db
}
