// router
// https://github.com/go-chi/chi/tree/master/_examples/todos-resource

// example - go-chi/docgen
// https://github.com/go-chi/chi/blob/master/_examples/rest/main.go

// flag
// https://golang.org/pkg/flag/

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"
)

var routes = flag.Bool("routes", false, "docgen - generated router documentation")
var docJSON = flag.Bool("json", false, "docgen - json")
var docMarkdown = flag.Bool("markdown", false, "docgen - markdown")

// go run . -routes -json
// go run . -routes -markdown
// go run . -routes -json -markdown

var docJSONFilename string = "chi-docgen-routes.json"
var docMarkdownFilename string = "chi-docgen-routes.md"

func main() {
	flag.Parse() // activate flag

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

	if *routes {
		if *docJSON {
			// Remove docJSONFilename if exists
			err := os.Remove(docJSONFilename)
			if err != nil {
				fmt.Println(err)
			}

			// Create new docJSONFilename
			f, err := os.Create(docJSONFilename)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			// Write JSONRoutesDoc to docJSONFilename
			text := docgen.JSONRoutesDoc(r)

			_, err = f.Write([]byte(text))
			if err != nil {
				fmt.Println(err)
			}
		}

		if *docMarkdown {
			// Remove docMarkdownFilename if exists
			err := os.Remove(docMarkdownFilename)
			if err != nil {
				fmt.Println(err)
			}

			// Create new docMarkdownFilename
			f, err := os.Create(docMarkdownFilename)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			// Write MarkdownRoutesDoc to docMarkdownFilename
			text := docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
				Intro: "Welcome to the generated router documentation by go-chi/docgen",
			})

			_, err = f.Write([]byte(text))
			if err != nil {
				fmt.Println(err)
			}
		}

		return
	}

	http.ListenAndServe(":3000", r)
}
