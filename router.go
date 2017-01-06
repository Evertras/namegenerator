package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// NewNamegenRouter returns a fully functioning API router for name generator
func NewNamegenRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/health", func(w http.ResponseWriter,
		r *http.Request,
		_ httprouter.Params) {
		fmt.Fprint(w, "Hi")
	})

	return router
}
