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
	fmt.Println("Iniciando o servidor Docker MySQL na porta 9010")
}

func GetDB() *gorm.DB {
	return db
}

// Aqui a primeira função (Connect()) serve para conectar ao banco de dados
// e a segunda função está disponibilizando o banco de dados pela variável "db"
