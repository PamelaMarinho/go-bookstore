package main

import (
	"fmt"
	"net/http"

	"github.com/PamelaMarinho/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println(http.ListenAndServe("localhost:8080", r))
}
