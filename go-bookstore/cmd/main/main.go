package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RonanMartin/Book-Management/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

	fmt.Println("Starting Docker MySQL server on port 9010")
}

// Here we are serving the application on the local port.
// we are also calling the route function that will handle the request.
