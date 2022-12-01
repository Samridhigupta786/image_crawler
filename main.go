package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handleRequest()

}
func handleRequest() {
	myRouter := mux.NewRouter()
	v1 := myRouter.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/submit", submitRequest)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
