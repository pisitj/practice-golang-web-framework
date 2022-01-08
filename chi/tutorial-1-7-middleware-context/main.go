// https://github.com/go-chi/chi#middleware-handlers

package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World."))
	})

	r.Group(func(r chi.Router) {
		r.Use(userMiddleware)
		r.Get("/user", getUserHandler)
	})

	http.ListenAndServe(":3000", r)
}

func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create new context from current request context
		// and assign key "user" to value of "123"
		ctx := context.WithValue(r.Context(), "user", "abc")

		// call next handler in the chain of middleware
		// passing the response writer and the updated request with new context value
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// read request context and get value of key "user"
	user := r.Context().Value("user").(string)

	// respond to client
	w.Write([]byte(fmt.Sprintf("Hi %v", user)))
}
