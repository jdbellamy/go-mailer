package main

import (
	"log"
	"net/http"
	"github.com/fatih/color"
	"github.com/uber-go/zap"
)

var logger = zap.New(zap.NewJSONEncoder())

func main() {
	port := ":8080"
	router := NewRouter()
	log.Printf("Listening on port %s\n", color.BlueString(port))
	log.Fatal(http.ListenAndServe(port, router))
}