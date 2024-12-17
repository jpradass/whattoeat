package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jpradass/whattoeat/db"
)

//go:embed templates/*.tmpl
var templates embed.FS

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFS(templates, "templates/*.tmpl"))
		recipe := db.GetRandomRecipe()

		tmpl.ExecuteTemplate(w, "index.tmpl", recipe)
	})

	r.Get("/favicon.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		http.ServeFile(w, r, "favicon-32x32.png")
	})

	http.ListenAndServe(":2550", r)
}
