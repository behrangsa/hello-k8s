package main

import (
	"github.com/behrangsa/hello-k8s/internal/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Info("Starting...")

	http.HandleFunc("/", handlers.GetIndex)
	http.HandleFunc("/hello", handlers.GetHello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}