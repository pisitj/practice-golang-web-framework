package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type articlesResource struct{}

func (rs articlesResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use(...)                 // apply routes Middleware

	r.Get("/", rs.List)    // GET /articles
	r.Post("/", rs.Create) // POST /articles

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(...)            // apply routes Middleware
		r.Get("/", rs.Get)       // GET /articles/{id}
		r.Put("/", rs.Update)    // PUT /articles/{id}
		r.Delete("/", rs.Delete) // DELETE /articles/{id}
	})

	return r
}

func (rs articlesResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of Articles."))
}

func (rs articlesResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a New Article."))
}

func (rs articlesResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a Article by ID."))
}

func (rs articlesResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a Article by ID."))
}

func (rs articlesResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a Article by ID."))
}
