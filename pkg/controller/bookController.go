package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/kelvin950/bookStoreManagement/pkg/models"
	"github.com/kelvin950/bookStoreManagement/pkg/utils"
)



func GETBooks(w http.ResponseWriter , r *http.Request){
   NewBook := models.GetAllBooks() 
   res , _ := json.Marshal(NewBook) 
   w.Header().Set("Content-Type" , "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(res)
}



func GetBookByID(w http.ResponseWriter , r *http.Request){

	params := mux.Vars(r)
 
	Id, err :=   strconv.ParseInt(params["bookid"],0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	} 
	NewBook  , _:=  models.GetBookById(Id)
  
    res ,  _ :=  json.Marshal(NewBook) 
	 
	w.Header().Set("Content_Type" , "application/json")
	w.WriteHeader(http.StatusAccepted) 
	w.Write(res)

}

func CreateBook(w http.ResponseWriter , r  *http.Request){
	 
	 createBook := &models.BOOK{}
	 utils.ParseBody(r , createBook) 
	 b:= createBook.CreateBook()

	 res , _ :=  json.Marshal(b) 
	
	 w.Header().Set("Content-Type" , "application/json" )
	w.WriteHeader(http.StatusAccepted) 
	
	w.Write(res)
}


func DeleteBook(w http.ResponseWriter , r *http.Request){

	 params :=  mux.Vars(r) 

	 Id,  err :=  strconv.ParseInt(params["bookid"] , 0,0) 
	 if err !=nil{
		fmt.Println("parsing error")
	 }

	 deletedBook := models.DeleteBook(Id) 

	 res, _ := json.Marshal(deletedBook) 
 

	 w.Header().Set("Content-Type" , "application/json")
	 w.WriteHeader(http.StatusAccepted)
	 w.Write(res)
	 
}


func UpdateBook(w http.ResponseWriter , r * http.Request){

	var updateBook =  &models.BOOK{} 

	utils.ParseBody(r,  updateBook) 
	  params:= mux.Vars(r) 

	 Id,  err :=  strconv.ParseInt(params["bookid"] , 0,0) 
	 if err !=nil{
		fmt.Println("parsing error")
	 }

	 book ,db := models.GetBookById(Id) 
	 
	  if updateBook.Name !=  ""{
		book.Name =  updateBook.Name
	  }
	  if updateBook.Author !=  ""{
		book.Author=  updateBook.Author
	  }
	  if updateBook.Publication !=  ""{
		book.Publication =  updateBook.Publication
	  }
      
	  db.Save(&book)
 
	  res , _ :=json.Marshal(book) 
	  
	  w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusAccepted)
	  w.Write(res)
	  
}