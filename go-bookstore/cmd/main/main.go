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

	fmt.Println("Iniciando o servidor Docker MySQL na porta 9010")
}

// Aqui estamos servindo a aplicação na porta local.
// também estamos chamando a função da rota que lidará com a requisição.
