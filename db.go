package main

import (
    "log"
    "os"
    "fmt"

    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
)


func Database() *gorm.DB {
    var login string = os.Getenv("FITQUESTDB")
    var db_string string = fmt.Sprintf("%s@tcp(fitquestdb:3306)/fitquest?charset=utf8&parseTime=True&loc=Local", login)

    db, err :=  gorm.Open("mysql", db_string)
    if err != nil {
        log.Fatal(err)
    }
    return db
}
