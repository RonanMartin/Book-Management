package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "Ronan:123teste/simplerest?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

// Aqui a primeira função (Connect()) serve para conectar ao banco de dados
// e a segunda função está disponibilizando o banco de dados pela variável "db"
