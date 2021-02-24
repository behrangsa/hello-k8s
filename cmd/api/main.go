package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Info("Starting...")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "I am Groot!")
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(writer, "Hello!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}