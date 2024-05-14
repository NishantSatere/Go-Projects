package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nishantdev/book-management-system/pkg/config"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:'name'"`
	Author      string `json:"author"`
	Publication string `json:"Publication"`
}

func init() {
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

// Functions to commincate with the database and used in the controller
func (b *Book) CreateBook() *Book {
	DB.NewRecord(b)
	DB.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	DB.Find(&books)
	return books
}

func GetBookById(id int) (*Book, *gorm.DB) {
	var getbook Book
	db := DB.Where("id = ?", id).Find(&getbook)
	return &getbook, db
}

func DeleteBook(id int) Book {
	var book Book
	DB.Where("id = ?", id).Delete(book)
	return book
}
