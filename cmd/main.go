package main

import (
	"log"
	"net/http"

	"github.com/goliiama/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

// crate the server/ localhost, recive route and send it to reoutes file which is bookstore-routes
func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
