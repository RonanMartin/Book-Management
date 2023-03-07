package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:test123@tcp(localhost:3307)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("Starting Docker MySQL server on port 9010")
}

func GetDB() *gorm.DB {
	return db
}

// Here the first function (Connect()) serves to connect to the database
// and the second function is making the database available through the "db" variable
