package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World."))
	})

	r.Post("/{id}", func(w http.ResponseWriter, r *http.Request) {
		// net/http - request
		// https://pkg.go.dev/net/http#Request
		fmt.Printf("host >> %v \n", r.Host)
		fmt.Printf("url >> %v \n", r.URL)
		fmt.Printf("query parameter >> %v \n", r.URL.Query())
		fmt.Printf("request body >> %v \n", r.Body)
		fmt.Printf("request header >> %v \n", r.Header)

		// chi - path parameter
		// https://github.com/go-chi/chi#url-parameters
		fmt.Printf("path parameter >> %v \n", chi.URLParam(r, "id"))

		// respond to the client
		w.Write([]byte("Hiiii - log request detail in terminal."))
	})

	http.ListenAndServe(":3000", r)
}
