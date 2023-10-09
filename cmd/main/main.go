package main

import (
	// "encoding/json"
	// "fmt"
	// "os"
	// "os/exec"
	// "sync"

	// _ "github.com/go-sql-driver/mysql"
	// _ "github.com/gorilla/mux"


"net/http"
"log"
"github.com/gorilla/mux"
_ "github.com/jinzhu/gorm"
"github.com/kelvin950/bookStoreManagement/pkg/routes"
)
// type pkg  struct{
// 	Dir  string   `json:"dir"`
//  ImportPath string  `json:"import_path"`
//  Name  string			`json:"name"`
//  Doc string		`json:"Doc"`
//  Target string   	`json:"target"`
//  Goroot  bool		`json:"goroot"`
//  Standard  bool   `json:"standard"`
//  Root string 		`json:"root"`
//  GoFiles []string		`json:"gofiles"`
//  Imports []string         `json:"imports"`
//  Deps []string		`json:"deps"`
// }
func main(){

	r :=  mux.NewRouter()

	routes.RegisterBookStoreRoutes(r) 
	http.Handle("/" , r)

	log.Fatal(http.ListenAndServe(":8000" ,  r))

// 	cmd :=  exec.Command("go" ,"list" , "-json" ,os.Args[1])

// 	var x pkg 
      
	 
//   output , err := cmd.Output() 
 
//    if err!=nil{
// 	panic(err)
//    }

//   json.Unmarshal(output , &x)
	
//  var wg sync.WaitGroup 

//  wg.Add(len(x.Deps))
//   for _ , r := range x.Deps  {

// 	fmt.Println(r)

// 	go func(r string){
        
// 		defer wg.Done()
// 		output , err :=  exec.Command("go" ,"list" , "-json" ,r).Output()

// 		if err!=nil{
// 			panic(err)
// 		}
//              println("=======")
// 		 println(string(output))
// 		 println("===========")
// 	}(r)
//   }

//   wg.Wait()
   

}