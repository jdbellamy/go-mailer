package main

import (
	"log"
	"net/http"
	"github.com/fatih/color"
)

func main() {
	port := ":8080"
	router := NewRouter()
	log.Panicf("Listening on port %s", color.BlueString(port))
	log.Fatal(http.ListenAndServe(port, router))
}

