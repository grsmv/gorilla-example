package controllers

import (
	"net/http"
	"fmt"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	if false {
		fmt.Fprintf(w, "never execute")
	} else {
		fmt.Fprintf(w, "root of all")
	}
}

func BooksHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Books Handler")
}
