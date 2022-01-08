// https://github.com/go-chi/chi/tree/master/_examples/todos-resource

package main

import (
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

	r.Mount("/users", usersResource{}.Routes())
	r.Mount("/articles", articlesResource{}.Routes())

	http.ListenAndServe(":3000", r)
}
