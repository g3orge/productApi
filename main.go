package main

import (
	"log"
	"os"

	"./handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	r := mux.NewRouter()
	r.HandleFunc("/product")
}
