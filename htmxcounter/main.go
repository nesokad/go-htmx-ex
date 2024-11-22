package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var counter = 0

func main() {
	fmt.Println("=== START ===")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			return
		}

		err = tmpl.ExecuteTemplate(w, "index.html", counter)
		if err != nil {
			return
		}
	})

	r.Post("/decrease", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("counter").Parse("{{ . }}"))

		counter -= 1

		tmpl.Execute(w, counter)
	})

	r.Post("/increase", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("counter").Parse("{{ . }}"))

		counter += 1

		tmpl.Execute(w, counter)
	})

	http.ListenAndServe(":3000", r)
}
