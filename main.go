package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewNamegenRouter()

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
