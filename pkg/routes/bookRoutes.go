package routes 
 
import (
	
	"github.com/gorilla/mux"
	"github.com/kelvin950/bookStoreManagement/pkg/controller"
)

var RegisterBookStoreRoutes =  func(r *mux.Router){

        r.HandleFunc("/books" , controller.CreateBook).Methods("POST")
		r.HandleFunc("/getBooks" , controller.GETBooks).Methods("GET")
		r.HandleFunc("/getBooks/{bookid}" , controller.GetBookByID).Methods("GET")
		r.HandleFunc("/books{bookid}" , controller.DeleteBook).Methods("DELETE")
		r.HandleFunc("/books{bookid}" , controller.UpdateBook).Methods("UPDATE")

}
