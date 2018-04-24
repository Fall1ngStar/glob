package main

import (
	"net/http"
	"log"
	"glob/routing"
)

func main() {
	log.Println("Starting server !")
	router := routing.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
