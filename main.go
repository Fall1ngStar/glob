package main

import (
	"net/http"
	"log"
	"glob/routing"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", routing.NewRouter()))
}
