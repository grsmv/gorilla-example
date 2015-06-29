package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/grsmv/gorilla-example/app"
)

func main() {
	r := mux.NewRouter()
	app.RegisterRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":9000", nil)
}