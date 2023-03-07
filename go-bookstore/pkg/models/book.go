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

// Here is calling functions from config/app.go to connect to the database and make it available.
// Makes the database information available inside the Struct Book.

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // db links us to the database, and NewRecord is a function that is part of the gorm package, which already has functions for requests
	db.Create(&b)   // we receive a Book from the database, create it inside the app and return
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books // Here the previous Find function loops inside the database and registers the books inside our books slice variable
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
