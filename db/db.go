package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/Thang-Tien/simple-car-booking/domains"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/cabrestapi")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&domains.Cab{}, &domains.Customer{}, &domains.Customerhistory{})
	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
