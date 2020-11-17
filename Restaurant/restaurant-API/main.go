package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"net/http"

	Model "restaurant/domain_model"
	DB "restaurant/persistence"

	"github.com/go-chi/chi"
)

func main() {
	port := "5000"
	r := registerRoutes()
	fmt.Println("Listening on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

//route declarations
func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/load", loadData)   //POST /load or /load?date=16050000
		r.Get("/buyers", getBuyers) //GET /buyers?first=50offset=100
		r.Get("/buyer", getBuyer)   //GET /buyer?id=234e2f
	})
	return r
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func loadData(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	timestampstring := r.URL.Query().Get("date")
	var timestamp int64
	if timestampstring == "" {
		timestamp = time.Now().Unix()
	} else {
		var err error
		timestamp, err = strconv.ParseInt(timestampstring, 10, 64)
		if err != nil {
			http.Error(w, "Invalid query parameters", 400)
			return
		}
	}

	exists := DB.QueryDateExists(timestamp)
	if exists == nil {
		http.Error(w, "Error while processing request", 500)
		return
	}
	if exists.Exists {
		json.NewEncoder(w).Encode(Model.Status{Success: false})
		return
	}

	status := DB.WriteBusinessDay(timestamp)
	if status == nil {
		http.Error(w, "Error while processing request", 500)
		return
	}
	json.NewEncoder(w).Encode(status)

}

func getBuyers(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	firstString := r.URL.Query().Get("first")
	offsetString := r.URL.Query().Get("offset")

	if firstString == "" || offsetString == "" {
		http.Error(w, "Invalid query parameters", 400)
		return
	}
	first, err := strconv.Atoi(firstString)
	if err != nil {
		http.Error(w, "Invalid query parameters", 400)
		return
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		http.Error(w, "Invalid query parameters", 400)
		return
	}

	buyers := DB.QueryBuyerList(first, offset)

	if buyers == nil {
		http.Error(w, "Error while processing request", 500)
		return
	}
	json.NewEncoder(w).Encode(buyers)

}

func getBuyer(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Invalid query parameters", 400)
		return
	}
	buyerdata := DB.QueryBuyerData(id)

	if buyerdata == nil {
		http.Error(w, "Error while processing request", 500)
		return
	}

	if buyerdata.Buyer.ID == "" {
		var empty struct{}
		json.NewEncoder(w).Encode(empty)
		return
	}

	json.NewEncoder(w).Encode(buyerdata)
}
