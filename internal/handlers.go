package internal

import (
	"fmt"
	"net/http"
)

func GetIndex(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "I am Groot!")
}

func GetHello(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Hello!")
}