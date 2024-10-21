package main

import (
	"fmt"
	"net/http"
	"service1/internal/http/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Printf("I'm alive! Hello, %s\n", name)
	})
    logMux := middleware.NewLogMux(mux)
	http.ListenAndServe(":8080", logMux)
}