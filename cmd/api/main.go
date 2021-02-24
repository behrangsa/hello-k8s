package main

import (
	"github.com/behrangsa/hello-k8s/internal"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Info("Starting...")

	http.HandleFunc("/", internal.GetIndex)
	http.HandleFunc("/hello", internal.GetHello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}