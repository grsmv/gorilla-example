package app

import "github.com/gorilla/mux"
import . "github.com/grsmv/gorilla-example/app/controllers"

func RegisterRoutes(r *mux.Router){
	r.HandleFunc("/",      RootHandler )
	r.HandleFunc("/books", BooksHandler)
}
