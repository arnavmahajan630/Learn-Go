package model

import (
	"github.com/OceanWhisperer/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

// init function for initializing
func init() {
	config.Connect()
	db = config.GetDB() // got the pointer to mysql DB
	db.AutoMigrate(&Book{})

}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var tb Book
	db := db.Where("Id=?", Id).Find(&tb)
	return &tb, db // returning the found book as well as the db pointer
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}

