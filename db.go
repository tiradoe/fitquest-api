package main

import (
    "log"

    "github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
)


func Database() *gorm.DB {
    db, err :=  gorm.Open("mysql", "admin:adminpass123@/fitquest?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        //panic("Failed to connect to database")
        log.Fatal(err)
    }
    return db
}
