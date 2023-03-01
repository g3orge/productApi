package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"gobred.com/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	r := mux.NewRouter()
	r.HandleFunc("/product")
}
