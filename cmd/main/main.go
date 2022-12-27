package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jesseinvent/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter();

	routes.RegisterBookStoreRoutes(r);

	http.Handle("/", r);

	fmt.Println("Server listening on http://localhost:9010");

	err := http.ListenAndServe("localhost:9010", r);

	if err != nil {
		log.Fatal(err);
	}
}