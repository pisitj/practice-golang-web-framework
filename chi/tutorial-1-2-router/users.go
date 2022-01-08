package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

type usersResource struct{}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use(...)                 // apply routes Middleware

	r.Get("/", rs.List)    // GET /users
	r.Post("/", rs.Create) // POST /users

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(...)            // apply routes Middleware
		r.Get("/", rs.Get)       // GET /users/{id}
		r.Put("/", rs.Update)    // PUT /users/{id}
		r.Delete("/", rs.Delete) // DELETE /users/{id}
	})

	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of Users."))
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a New User."))
}

func (rs usersResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a User by ID."))
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a User by ID."))
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a User by ID."))
}
