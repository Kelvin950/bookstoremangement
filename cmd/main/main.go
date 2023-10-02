package main

import (
	

	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	"github.com/kelvin950/bookStoreManagement/pkg/routes"
)

func main(){

	r :=  mux.NewRouter()

	routes.RegisterBookStoreRoutes(r) 
	http.Handle("/" , r)

	log.Fatal(http.ListenAndServe(":8000" ,  r))
}