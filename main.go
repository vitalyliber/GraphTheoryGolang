package main

import (
	"./handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/create", handlers.CreateVertex)
	http.ListenAndServe(":8080", nil)
}
