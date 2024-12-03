package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Recipe struct {
	Id    string
	Title string
	Image string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("tpls/index.html"))
		recipe := map[string]Recipe{
			"Recipe": {
				Title: "Spaghetti",
				Id:    "1291-41042-44214",
				Image: "https://plus.unsplash.com/premium_photo-1677000666741-17c3c57139a2?q=80&w=2787&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			},
		}

		tmpl.Execute(w, recipe)
	})

	// db.GetConn()
	http.ListenAndServe(":2550", r)
}
