package models 

import (
	"github.com/jinzhu/gorm" 
	"github.com/kelvin950/bookStoreManagement/pkg/config"
	
)


var db *gorm.DB 

type BOOK struct {

	gorm.Model
	Name string `gorm:""json:"name"`
	Author string 	`json:"author"`
	Publication string `json:"publication"`
}

func init(){

	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&BOOK{})
}

func (b *BOOK)CreateBook() *BOOK{

 
	db.NewRecord(b) 
	db.Create(&b) 


	return b
}

func GetAllBooks() []BOOK{

	var books []BOOK
	db.Find(&books)

	return books
}

func GetBookById(id int64) (*BOOK , *gorm.DB){
	var getBook BOOK 
	db := db.Where("ID=?" ,id ).Find(&getBook) 

	return &getBook , db
}

func DeleteBook(id int64)BOOK{

	var book BOOK 
	db.Where("ID=?" , id).Delete(book) 

	return book
}