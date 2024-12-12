package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jpradass/whattoeat/db"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		recipe := db.GetRandomRecipe()
		// recipe := map[string]Recipe{
		// 	"Recipe": {
		// 		Title:       "Spaghetti",
		// 		Id:          "1291-41042-44214",
		// 		Image:       "https://plus.unsplash.com/premium_photo-1677000666741-17c3c57139a2?q=80&w=2787&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		// 		Description: "Just a regular dish of spaghetti",
		// 		Ingredients: []Ingredient{
		// 			{"🍝 400g spaghetti"},
		// 			{"🥓 200g pancetta, diced"},
		// 		},
		// 		Steps: []Step{
		// 			{"Boil 1L of water in a pot"},
		// 			{"Cut the vegetables and prepare everything"},
		// 			{"Put a pan in the fire and put the vegetables on it"},
		// 		},
		// 	},
		// }

		tmpl.Execute(w, recipe)
	})

	r.Get("/favicon.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		http.ServeFile(w, r, "favicon-32x32.png")
	})

	// db.GetConn()
	http.ListenAndServe(":2550", r)
}
