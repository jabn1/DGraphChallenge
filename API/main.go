package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/load", loadData).Query().Get("date") //POST /load?date=16050000
		r.Get("/{buyers}", getBuyers)                 //GET /buyers
		r.Get("/buyer", getBuyer).Query().Get("id")   //GET /buyer?id=234e2f
	})
	return r
}
