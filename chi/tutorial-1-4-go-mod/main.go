// https://github.com/go-chi/chi/tree/master/_examples/limits

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// chech Timeout
	r.Group(func(r chi.Router) {
		// Stop processing after 2.5 seconds.
		r.Use(middleware.Timeout(2500 * time.Millisecond))

		r.Get("/random", func(w http.ResponseWriter, r *http.Request) {
			rand.Seed(time.Now().Unix())

			// Processing will randomly take 1-5 seconds.
			randomProcessTime := time.Duration(rand.Intn(4)+1) * time.Second
			fmt.Printf("Process Time - %v seconds\n", randomProcessTime)

			select {
			case <-r.Context().Done():
				w.WriteHeader(http.StatusGatewayTimeout)
				w.Write([]byte(fmt.Sprintf("Fail to Process in %v seconds\n", randomProcessTime)))
				return

			case <-time.After(randomProcessTime): // This channel simulates some hard work.
			}

			w.Write([]byte(fmt.Sprintf("Success to Processed in %v seconds\n", randomProcessTime)))
		})
	})

	// check Throttle
	r.Group(func(r chi.Router) {
		// Only one request will be processed at a time.
		r.Use(middleware.Throttle(1))

		r.Get("/submit", func(w http.ResponseWriter, r *http.Request) {
			<-time.After(5 * time.Second) // This channel simulates some hard work.
			w.Write([]byte("Processed\n"))
		})

	})

	http.ListenAndServe(":3000", r)

}
