package models

import (
	"github.com/RonanMartin/Book-Management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Aqui está chamando as funções do config/app.go para se conectar ao banco de dados e disponibilizá-lo.
// parece disponibilizar as informações do banco dedados dentro do Struct Book.

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // db nos liga ao banco de dados, e NewRecord é uma função que faz parte do package gorm, que já possui funções para requisições
	db.Create(&b)   // recebemos um Book do banco de dados, criamos ele dentro do app e retornamos
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books // Aqui a função prévia Find, passa um loop dentro do banco de dados e registra os livros dentro da nossa váriável slice de books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
